package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"polychem-auth/config"
	"polychem-auth/models"
	"polychem-auth/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// Login - ورود با یوزرنیم و پسورد
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: "Username and password are required",
		})
		return
	}

	// دریافت IP و User-Agent
	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()

	// جستجوی کاربر
	var admin models.Admin
	err := h.DB.QueryRow(
		"SELECT id, username, email, password_hash, phone, full_name, is_active FROM admins WHERE username = ?",
		req.Username,
	).Scan(
		&admin.ID, &admin.Username, &admin.Email, &admin.PasswordHash,
		&admin.Phone, &admin.FullName, &admin.IsActive,
	)

	if err == sql.ErrNoRows {
		h.logLoginAttempt(nil, req.Username, "login_failed", ipAddress, userAgent, false, "user_not_found")
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid username or password",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "database_error",
		})
		return
	}

	// بررسی فعال بودن حساب
	if !admin.IsActive {
		h.logLoginAttempt(&admin.ID, req.Username, "login_failed", ipAddress, userAgent, false, "account_disabled")
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Error:   "account_disabled",
			Message: "Your account has been disabled",
		})
		return
	}

	// بررسی پسورد
	if !utils.CheckPasswordHash(req.Password, admin.PasswordHash) {
		h.logLoginAttempt(&admin.ID, req.Username, "login_failed", ipAddress, userAgent, false, "wrong_password")
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid username or password",
		})
		return
	}

	// تولید توکن‌ها
	token, err := utils.GenerateToken(admin.ID, admin.Username, admin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "token_generation_failed",
		})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(admin.ID, admin.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "token_generation_failed",
		})
		return
	}

	// ذخیره Session
	expiresAt := time.Now().Add(time.Duration(config.AppConfig.JWTExpiryHours) * time.Hour)
	_, err = h.DB.Exec(
		`INSERT INTO sessions (user_id, token, refresh_token, expires_at, ip_address, user_agent)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		admin.ID, token, refreshToken, expiresAt, ipAddress, userAgent,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "session_creation_failed",
		})
		return
	}

	// به‌روزرسانی زمان آخرین ورود
	_, _ = h.DB.Exec("UPDATE admins SET last_login = NOW() WHERE id = ?", admin.ID)

	// ثبت لاگ موفق
	h.logLoginAttempt(&admin.ID, req.Username, "login_success", ipAddress, userAgent, true, "")

	// پاسخ
	c.JSON(http.StatusOK, models.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User: models.AdminPublic{
			ID:       admin.ID,
			Username: admin.Username,
			Email:    admin.Email,
			Phone:    admin.Phone,
			FullName: admin.FullName,
		},
	})
}

// Logout - خروج از سیستم
func (h *AuthHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
		return
	}

	// حذف Bearer از ابتدای توکن
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// حذف Session
	_, _ = h.DB.Exec("DELETE FROM sessions WHERE token = ?", token)

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

// ثبت تلاش برای ورود
func (h *AuthHandler) logLoginAttempt(userID *int, username, action, ip, ua string, success bool, reason string) {
	_, _ = h.DB.Exec(
		`INSERT INTO login_logs (user_id, username, action, ip_address, user_agent, success, failure_reason)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		userID, username, action, ip, ua, success, reason,
	)
}