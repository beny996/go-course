package main

import (
	"errors"
	"fmt"
)

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func addSignature(message string) string {
	return message + " Kind regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}

func test(messages []string, formatter func(string) string) {
	formattedMessages := getFormattedMessages(messages, formatter)
	fmt.Println("Here are the formatted messages : ")
	for _, message := range formattedMessages {
		fmt.Println(message)
	}
}

// currying

// function currying is the practice of writing a function that takes a function (or more) as input, and returns a new function

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a string, b string) {
		fmt.Println(formatter(a, b))
	}
}

func test2(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("==================")
	logger := getLogger(formatter)

	fmt.Println("Logs :")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func coloDelimit(first, second string) string {
	return first + ": " + second
}

func commaDelimit(first, second string) string {
	return first + ", " + second
}

// defer

// defer keyword allows a function to be executed just before its enclosing function returns (in the end)

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	user, ok := users[name]

	defer delete(users, name)

	if !ok {
		return logNotFound
	}

	if user.admin {
		return logAdmin
	}
	return logDeleted

}

type user struct {
	name   string
	number int
	admin  bool
}

func test3(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("=============")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}

// closures

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

type emailBill struct {
	costInPennies int
}

func test4(bills []emailBill) {
	defer fmt.Println("====================")
	countAdder, costAdder := adder(), adder()

	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

// anonymous functions

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(msg string) int {
			return len(msg) * 2
		}, message)
	}
}

func test5(messages []string) {
	defer fmt.Println("=======================")
	printReports(messages)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf("Message : '%s' Cost : %v cents\n", message, cost)
}

func main() {

	test([]string{
		"Test0",
		"Test1",
		"Test2",
		"Test3",
		"Test4",
		"Test5",
	}, addSignature)

	test([]string{
		"Test0",
		"Test1",
		"Test2",
		"Test3",
		"Test4",
		"Test5",
	}, addGreeting)

	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}

	test2("Error on database server", dbErrors, coloDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}

	test2("Error on mail server", mailErrors, commaDelimit)

	users := map[string]user{
		"john": {
			name:   "john",
			number: 123123123,
			admin:  true,
		},
		"elon": {
			name:   "elon",
			number: 1231244546,
			admin:  true,
		},
		"jack": {
			name:   "jack",
			number: 12344244546,
			admin:  false,
		},
	}

	test3(users, "john")
	test3(users, "elon")
	fmt.Println("Final users :")
	for _, user := range users {
		fmt.Println(user.name)
	}

	test4([]emailBill{
		{43},
		{53},
		{53},
		{12},
		{78},
		{32},
		{23},
	})

	test5([]string{
		"This is a first test message",
		"This is a second test message, that will increase the cost",
		"This is a third test message. This will be the longest and the most expensive one",
	})
}
