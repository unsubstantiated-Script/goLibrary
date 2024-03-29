package application_basics

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func RollBCrypt() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	logPwrd1 := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(logPwrd1))
	if err != nil {
		fmt.Println("You can't login")
		return
	}

	fmt.Println("You're logged in")
}
