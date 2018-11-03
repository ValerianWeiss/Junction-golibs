package jwt

import (
	"errors"
	"os"
	"strconv"

	"github.com/SermoDigital/jose/jws"
)

// GetClaim returns a claim from a jwt as a string.
// the token should contianer the "Bearer" prefix
func GetClaim(claim string) (string, error) {
	token, err := getToken()

	if err != nil {
		return "", err
	}

	jwt, err := jws.ParseJWT([]byte(token))

	if err != nil {
		return "", err
	}

	claims := jwt.Claims()
	value := claims.Get(claim)

	switch value.(type) {
	case string:
		return value.(string), nil
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', 2, 64), nil
	default:
		return "", errors.New("The claimvalue has an invalid type")
	}
}

func getToken() (string, error) {
	bearer := os.Getenv("Http_Authorization")

	if bearer == "" {
		return "", errors.New("No token in header")
	}

	// Cutting of the Bearer prefix
	token := bearer[7:]
	return token, nil
}

// GetClaimSub returns the sub claim of a JWT
func GetClaimSub() (string, error) {
	return GetClaim("sub")
}

// GetUserid returns the userid from the JWT token. OpenFaaS stores the
// "Authorization" as Environment variable under "Http_Authorization"
func GetUserid() (string, error) {
	auth0id, err := GetClaimSub()

	// The auth0id looks like: auth0|5bcc993740asfeae771f
	// cutting off the prefix
	userid := auth0id[6:]
	return userid, err
}
