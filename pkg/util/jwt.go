package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/pkg/setting"
)

var jwtSecret []byte

type Claims struct {
	jwt.StandardClaims
	Identity *models.User `json:"identity"`
}

// GenerateToken generate tokens used for auth
func GenerateToken(user *models.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    setting.Config.Address,
		},
		Identity: &models.User{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		},
	}

	accessTokenString, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString([]byte(setting.Config.JWTSigningKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.Config.JWTVerificationKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
