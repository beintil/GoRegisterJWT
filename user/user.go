package user

import (
	"golang.org/x/crypto/bcrypt" //пакет bcrypt используется для шифрования/расшифровки паролей.
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// 'хеширует' пароль
func (user *User) HashPasswod(password string) error {
	hashPas, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return err
	}

	user.Password = string(hashPas)

	return nil
}

// сравниваем хешированный пароль с его введенным аналогом возращая nil при совподении
func (user *User) ComparesPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
