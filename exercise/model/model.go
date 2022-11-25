package model

import (
	"fmt"
)

type NormalUser struct {
	Name string
	Account
}

type VIPUser struct {
	Name      string
	Privilege string
	Account
}

type Account struct {
	Username string
	Password string
}

func (account Account) CheckPassword(password string) bool {
	defer func() {
		fmt.Println("Password Checked")
	}()
	return account.Password == password
}

func (normalUser NormalUser) UserType() string {
	return "Normal User"
}

func (vipUser VIPUser) UserType() string {
	return "VIP"
}
