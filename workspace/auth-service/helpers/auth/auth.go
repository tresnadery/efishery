package helpers

import(
	"fmt"
	"strings"
	"auth-service/domain"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetClaimsJWT(tokenString string) (map[string]interface{}, error) {
	if tokenString == "" {
		return nil, domain.ErrTokenNotFound
	}
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		hmacSecretString := "secret"
		hmacSecret := []byte(hmacSecretString)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(domain.ErrUnexpectedSingningMethod.Error(), token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// return claims["user_id"].(string), nil
		return map[string]interface{}{
			"phone_number" : claims["phone_number"].(string),
			"name" : claims["name"].(string),
			"role_name": claims["role_name"].(string),
			"created_at": claims["created_at"].(string),
		}, nil
	}
	return nil, err
}

func HashAndSaltPassword(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {		
		return "", err
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}