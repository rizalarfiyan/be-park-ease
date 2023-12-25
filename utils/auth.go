package utils

import (
	"be-park-ease/config"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashString(str string) string {
	conf := config.Get()
	h := sha1.New()
	_, err := h.Write([]byte(str + conf.Auth.TokenSalt))
	if err != nil {
		return str
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func GenerateToken(userId int32) string {
	uniq := time.Now().UnixMicro()
	return HashString(fmt.Sprint(userId, uniq))
}
