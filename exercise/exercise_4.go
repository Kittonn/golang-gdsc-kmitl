package exercise

import (
	"fmt"
	"golang-gdsc-kmitl/exercise/model"
)

type User interface {
	UserType() string
}

func printType(user User) {
	println(user.UserType())
}

func Exercise_4() {
	normalUser := model.NormalUser{
		Name: "kittipod",
		Account: model.Account{
			Username: "k1tt1p0d",
			Password: "qwerty5678",
		},
	}

	vipUser := model.VIPUser{
		Name:      "Kitton",
		Privilege: "VIP",
		Account: model.Account{
			Username: "k1tt0n",
			Password: "qwerty5678",
		},
	}
	printType(normalUser)
	printType(vipUser)
	fmt.Println(normalUser.CheckPassword("qwerty5678"))
	fmt.Println(vipUser.CheckPassword("qwerty5678"))
	fmt.Println("-----")
}
