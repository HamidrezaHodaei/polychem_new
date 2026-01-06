package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Enter username: ")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)
    
    fmt.Print("Enter email: ")
    email, _ := reader.ReadString('\n')
    email = strings.TrimSpace(email)
    
    fmt.Print("Enter phone (e.g. 09123456789): ")
    phone, _ := reader.ReadString('\n')
    phone = strings.TrimSpace(phone)
    
    fmt.Print("Enter full name: ")
    fullName, _ := reader.ReadString('\n')
    fullName = strings.TrimSpace(fullName)
    
    fmt.Print("Enter password: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)
    
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
    fmt.Println("║                    ADMIN CREATED SUCCESSFULLY                  ║")
    fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
    fmt.Printf("║ Username:  %-51s║\n", username)
    fmt.Printf("║ Email:     %-51s║\n", email)
    fmt.Printf("║ Phone:     %-51s║\n", phone)
    fmt.Printf("║ Full Name: %-51s║\n", fullName)
    fmt.Printf("║ Password:  %-51s║\n", password)
    fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
    fmt.Println("║                   COPY THIS SQL COMMAND:                       ║")
    fmt.Println("╚═══════════════════════════════════════════════════════════════╝\n")
    
    fmt.Printf("INSERT INTO admins (username, email, password_hash, phone, full_name)\n")
    fmt.Printf("VALUES ('%s', '%s', '%s', '%s', '%s');\n\n",
        username, email, string(hash), phone, fullName)
}