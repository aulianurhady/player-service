package jwt

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type M map[string]interface{}

type MyClaims struct {
	jwt.StandardClaims
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	RegisterAt string `json:"register_at"`
}

var (
	APPLICATION_NAME          = "Player Service"
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("the secret of player service")
)

func CreateClaimJwt(id uint, username, registerAt string) string {

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Id:         id,
		Username:   username,
		RegisterAt: registerAt,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return ""
	}

	return signedToken
}

func ClaimToken(tokenString string) (id string, username string, registerAt string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", "", errors.New("token not valid")
	}

	return claim["id"].(string), claim["username"].(string), claim["register_at"].(string), nil
}
