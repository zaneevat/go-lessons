package main

import "fmt"

var _ User = &superUser{}

type superUser struct {
	isBlocked bool
}

func (s superUser) Block() {
	s.isBlocked = true
}

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
	isBlocked           bool
}

func (u user) Block() {
	u.isBlocked = true
}

type User interface {
	Block()
}

func NewUser(fio, address, phone string) User {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
	return &u
}

func main() {
	u := NewUser("", "", "")
	u.Block()
	fmt.Println(u)
}
