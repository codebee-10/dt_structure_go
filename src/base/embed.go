package main

import "fmt"

type notifier interface {
	notify()
}

type user struct{
	name string
	email string
}

type admin struct {
	user
	level string
}

func (u *user) notify(){
	fmt.Println("user:", u.name, u.email)
}

func (u *admin) notify(){
	fmt.Println("admin:",u.name, u.email)
}

func sendNotifiation(n notifier){
	n.notify()
}

func main() {
	ad := admin{
		user: user{
			name: "Ethan",
			email: "roancsu@163.com",
		},
		level: "super",
	}

	sendNotifiation(&ad)
	ad.user.notify()
	ad.notify()
}

