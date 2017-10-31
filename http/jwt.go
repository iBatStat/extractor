package http

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	privateKey, publicKey []byte
)

func init() {
	privateKey, _ = ioutil.ReadFile(os.Getenv("JWT_PRIVATE_KEY"))
	publicKey, _ = ioutil.ReadFile(os.Getenv("JWT_PUBLIC_KEY"))
}

// we need functions that do the following
//
// 1. generate a new jwt token and stick user claims
//    user claims can be - user id and device type for now
//    for now - expireAt will be 1 hours -- default and hard coded
//    we will also stick the genereated token in the header of the request and redirect to desired site
// 2. authenticate a jwt on a http reqeust
// 3. possibly we want to re-validate a jwt with expire at extended

func generateNew(userId, phoneType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":      userId,
		"phoneType": phoneType,
		"exp":       time.Now().Add(1 * time.Hour),
	})
	tokenstr, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenstr, nil
}

func authenticate(r *http.Request) bool {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	return true
}
