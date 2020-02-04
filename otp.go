package main

import (
	"fmt"
	"github.com/creachadair/otp"
	"os"
)

func Codes(activationCode string) (token string, _ error) {
	var conf otp.Config
	if err := conf.ParseKey(activationCode); err != nil {
		return "", err
	}
	token = conf.TOTP()
	return
}

func main() {
	var activationCode string
	if len(os.Args) <= 1 {
		fmt.Println("Missing activationCode...")
	} else {
		//Get the default first output as the activation Code
		activationCode = os.Args[1]

		token, e := Codes(activationCode)
		if e != nil {
			panic(e)
		}
		fmt.Println(token)
	}
}
