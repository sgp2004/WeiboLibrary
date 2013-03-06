package service

import (
	"../dao"
	"os/exec"
)

func Auth(user, pwd, pwdOriginal string) bool {
	userInfo := dao.GetRegisterUser(user, pwd)
	if userInfo != nil {
		return true
	}
	print("check user by office code----")
	cmd := exec.Command("java", "LDAPUtil", user, pwdOriginal)
	out, err := cmd.Output()
	print(string(out))
	if err != nil {
		println("error:", err.Error())
		return false
	}
	if string(out) == "true" {
		return true
	}

	return false
}
