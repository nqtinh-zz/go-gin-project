package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        int32     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password,omitempty" db:"password"`
	Accounts  []Account `json:"accounts,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at"`
}

type CreateUserReq struct {
	Username string `json:"username" db:"username" validate:"required,alphanum"`
	Password string `json:"password" db:"password" validate:"required"`
}

type LoginReq struct {
	Username string `json:"username" db:"username" validate:"required,alphanum"`
	Password string `json:"password" db:"password" validate:"required"`
}

type LoginResp struct {
	AccessToken string `json:"access_token,omitempty"`
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	Identity *User `json:"identity"`
}
