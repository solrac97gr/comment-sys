package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/solrac97gr/comments/models"
)

var (
	privateKey *rsa.PrivateKey
	// PublicKey use for validate the token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal(err)
	}

	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Can not parse the private key")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Can not parse the public key")
	}
}

// GenerateJWT generate the token for the client
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer: "Carlos García",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("can not signed the token")
	}
	return result
}
