package main

import (
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

func removeProfanity(message *string) {
	if message == nil {
		return
	}

	messageValue := *message
	messageValue = strings.ReplaceAll(messageValue, "dang", "****")
	messageValue = strings.ReplaceAll(messageValue, "shoot", "*****")
	messageValue = strings.ReplaceAll(messageValue, "heck", "****")
	*message = messageValue

}

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

// zero value of a pointer is nil (if the pointers doesn't point to anything, it's value will be nil)

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test2(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage(newMessage)
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("===================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func main() {

	sendMessage(Message{
		Recipient: "Vukasin",
		Text:      "This is the first test message",
	})

	sendMessage(Message{
		Recipient: "Lane",
		Text:      "This is the second test message",
	})

	test([]string{
		"dang this is bad",
		"heck this is bad",
		"shoot this is bad",
	})

	test2(&email{
		message:     "this is my first draft",
		fromAddress: "from@mail.com",
		toAddress:   "to@email.com",
	}, "this is my second draft")

	test2(&email{
		message:     "this is my first draft",
		fromAddress: "from@mail.com",
		toAddress:   "to@email.com",
	}, "this is my third draft")

	test2(&email{
		message:     "this is my first draft",
		fromAddress: "from@mail.com",
		toAddress:   "to@email.com",
	}, "this is my fourth draft")
}
