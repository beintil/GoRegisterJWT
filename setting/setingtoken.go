package setting

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// что то типо ключа для создания токена jwt
var jwtKey = []byte("gentykey")

type JWTstruct struct {
	Username string `json:"username"`
	Email    string `json:"email"`

	// Определенные данные, которые соответствуют стандарту jwt
	jwt.StandardClaims
}

func GenerateJWT(email string, username string) (tokenKey string, err error) {
	//Устанавливаем время действия токена, по истечению которого он становится не доступен
	endTime := time.Now().Add(24 * time.Hour)

	// здесь мы устанавливаем хароктеристики нашего токена(время действия) и 'шифруем' в нем наши данные
	generate := &JWTstruct{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: endTime.Unix(),
		},
	}

	// создаем токен типа HS256б передавая в него все характеристики
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, generate)
	tokenKey, err = token.SignedString(jwtKey)
	return
}

//
func CheckToken(signedToken string) (err error) {
	// проводим чтото типо анализа токена
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTstruct{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	// проверяем, смогли ли мы проанализировать токен
	chec, ok := token.Claims.(*JWTstruct)
	if !ok {
		err = errors.New("analyssis error")
		return
	}

	// проверяем, истекло ли время действия созданного токена
	if chec.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("the token has expired")
		return
	}

	return
}
