package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Password encryption error. Error: " + err.Error())
	}

	return string(hashedPwd)
}

func CheckPassword(pwdFromDb string, pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pwdFromDb), []byte(pwd))

	return err
}
