package main

import (
	"fmt"
)

func main() {
	var tom Person
	tom.name = "tom"
	tom.eat()
	var user User
	user.name = "tom"
	user.passwd = "123456"
	user.login(user)
}

type Person struct {
	name string
}

func (per Person) eat() {
	fmt.Printf("name: %v eat...\n", per.name)
}

type User struct {
	name   string
	passwd string
}

func (user User) login(loginUser User) bool {
	if loginUser.name == "tom" && loginUser.passwd == "123456" {
		fmt.Printf("name:%v login success\n", loginUser.name)
		return true
	} else {
		fmt.Printf("name:%v login fail\n", loginUser.name)
		return false
	}

}
