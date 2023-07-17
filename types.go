package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type LoginAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateAccountRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	Password  string    `json:"-"`
}

func NewAccount(email, name, passwd string) *Account {
	encPasswd := CreateHash(passwd)
	return &Account{
		email,
		name,
		time.Now(),
		encPasswd,
	}
}

func CreateHash(s string) string {
	hash := sha256.New()

	return hex.EncodeToString(hash.Sum([]byte(s)))
}
