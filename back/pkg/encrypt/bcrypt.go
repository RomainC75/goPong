package encrypt

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error){
	fmt.Println("PASSWORD to encrypt : ", pwd)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPwd string, receivedPwd string) error{
	fmt.Println(hashedPwd,"_______", receivedPwd)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(receivedPwd))
	if err != nil {
		fmt.Println("==> bcrypt error : ", err)
		return err
	}
	return nil
}
