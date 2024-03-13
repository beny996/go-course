package main

import (
	"fmt"
)

type messageToSend struct {
	phoneNumber int
	message     string
}

//embedded structs

type messageToSend2 struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

type sender struct {
	rateLimit int
	user
}

//methods on structs

type rect struct {
	width  int
	height int
}

//this creates a method on struct "rect"
func (r rect) area() int {
	return r.width * r.height
}

func test(m messageToSend) {
	fmt.Printf("sending message: '%s' to: %v\n", m.message, m.phoneNumber)
}

func test2(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rate limit:", s.rateLimit)
}

func canSendMessage(mToSend messageToSend2) bool {
	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.recipient.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}

type car struct {
	make  string
	model string
}

type truck struct {
	//"car" is embedded, so the definition of a "truck" now also additionally contains all of the fields of the car struck

	car
	bedSize int
}

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
}

func test3(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
}

func main() {

	//one more way of defining a struct is when we don't define a type for the struct first, but we use an anonymous struct type when defining it

	myCar := struct {
		make  string
		model string
	}{
		make:  "Lamborghini",
		model: "Aventador",
	}

	fmt.Println(myCar.make)

	test(messageToSend{
		phoneNumber: 15454554,
		message:     "this is a text message",
	})

	truck := truck{
		bedSize: 10,
		car: car{
			make:  "Ford",
			model: "Raptor",
		},
	}

	fmt.Println(truck)

	test2(sender{
		rateLimit: 10000,
		user: user{
			name:   "Vukasin",
			number: 1231231,
		},
	})

	r := rect{
		width:  10,
		height: 5,
	}

	fmt.Println(r.area())

	test3(authenticationInfo{
		username: "Vukasin",
		password: "Test password",
	})
}
