package middleware

import (
	"database/sql"
	"net/http"
	"strings"

	"polychem-auth/models"
	"polychem-auth/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware - بررسی توکن JWT
func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "unauthorized",
				Message: "Authorization header required",
			})
			c.Abort()
			return
		}

		// حذف Bearer از ابتدای توکن
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "invalid_token_format",
				Message: "Token must start with 'Bearer '",
			})
			c.Abort()
			return
		}

		// اعتبارسنجی توکن
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "invalid_token",
				Message: "Token is invalid or expired",
			})
			c.Abort()
			return
		}

		// بررسی وجود Session در دیتابیس
		var sessionID int
		err = db.QueryRow(
			"SELECT id FROM sessions WHERE token = ? AND expires_at > NOW()",
			tokenString,
		).Scan(&sessionID)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error:   "session_expired",
				Message: "Session has expired or been revoked",
			})
			c.Abort()
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "database_error",
			})
			c.Abort()
			return
		}

		// ذخیره اطلاعات کاربر در Context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		c.Next()
	}
}