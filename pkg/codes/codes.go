package codes

import "github.com/creachadair/otp"

func Codes(activationCode string) (token string, _ error) {
	var conf otp.Config
	if err := conf.ParseKey(activationCode); err != nil {
		return "", err
	}
	token = conf.TOTP()
	return
}
