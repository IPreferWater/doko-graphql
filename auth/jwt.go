package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("secret")
)

func CreateToken(userid int) (string, error) {

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString(SecretKey)
}

func VerifyToken(authorisation string) (bool, error) {
	log.Info(authorisation)
	return true, nil
}
