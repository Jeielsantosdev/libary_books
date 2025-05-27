package services


import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

// CheckPasswordHash verifica se a senha fornecida corresponde ao hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        log.Println("Erro ao verificar senha:", err)
        return false
    }
    return true
}