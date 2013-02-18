package dao

import "fmt"
import "testing"

func TestCreateUser(t *testing.T) {
	user := new(User)
	user.UserName = "daodao"
	user.Name = "好人"
	fmt.Println(user)
	CreateUser(user)
}
func TestQueryUserByUserName(t *testing.T) {
	user := QueryUserByUserName("daodao")
	fmt.Println("user = \n")
	fmt.Println(user)
}
func TestUpdateCount(t *testing.T) {

	user := QueryUserByUserName("daodao")
	UpdateCount(1, user.Id, 1)
	UpdateCount(1, user.Id, 0)
	user = QueryUserByUserName("daodao")
	fmt.Println("user after update= \n")
	fmt.Println(user)
}
