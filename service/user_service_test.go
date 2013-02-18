package service

import "fmt"
import "testing"

func TestAuth(t *testing.T) {
	user := "guanpu"
	pwd := "kjjnhi6j"
	fmt.Println(pwd == string("kjjnhi6j"))
	if Auth(user, pwd) {
		fmt.Println("auth success")
	} else {
		fmt.Println("auth false")
	}
}
