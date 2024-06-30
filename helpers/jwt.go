package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	SECRET_KEY []byte = []byte("20242213960490267598567908189775563674735941821979632997103732363925894510946901202218157")
)

func HasPass(password string) (string, error) {
	salt := 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUserJWT(id uint64, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
		})
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateUserJWT(token string) (bool, any) {
	jwttoken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if err != nil {
		return false, nil
	}

	if _, ok := jwttoken.Claims.(jwt.MapClaims); !ok && !jwttoken.Valid {
		return false, nil
	}
	return jwttoken.Valid, jwttoken.Claims.(jwt.MapClaims)
}
