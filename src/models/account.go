package models

import (
	"golang.org/x/crypto/bcrypt"
)

// Account table
type Account struct {
	ID       uint   `gorm:"primary_key" json:"id" form:"id"`
	Username string `gorm:"not null;unique;index" json:"username" form:"username"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Active   uint   `gorm:"default:1"`
}

// HasPassword encrypt password
// https://gowebexamples.com/password-hashing/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compare plain password w hash password
// https://gowebexamples.com/password-hashing/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
