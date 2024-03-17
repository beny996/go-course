package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, your birthday is on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your '%s' report is ready. You've sent %v messages", sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}

	return 0.01 * float64(len(e.body))
}

func (sms sms) cost() float64 {
	if !sms.isSubscribed {
		return 0.05 * float64(len(sms.body))
	}

	return 0.01 * float64(len(sms.body))
}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func print(p printer) {
	p.print()
}

func test1(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ... \n", e.cost())
	p.print()
}

type Copier1 interface {
	//this is bad because we don't know anything about the parameters and the return value
	Copy(string string) int
}

//this is better because we now know what are those parameters that we receive (syntax is different, we now have comma separating two parameters)
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

type invalid struct{}

func (i invalid) cost() float64 {
	return 0.0
}

func getExpenseReport(e expense) (string, float64) {
	//this checks if e is an instance of email, if it is, ok will be true, if it is not, it will be false
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}

	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}
	return "", 0.0
}

func test2(e expense) {
	address, cost := getExpenseReport(e)

	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost $%.2f \n", address, cost)

	case sms:
		fmt.Printf("Report: The sms going to %s will cost $%.2f \n", address, cost)

	default:
		fmt.Println("Report: Invalid report!")
	}
}

//TYPE SWITCHES

func getExpenseReport2(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()

	case sms:
		return v.toPhoneNumber, v.cost()

	default:
		return "", 0.0
	}

	//this is a type switch

}

func main() {

	test(sendingReport{
		reportName:    "First report",
		numberOfSends: 25,
	})

	test(birthdayMessage{
		birthdayTime:  time.Date(1996, 04, 21, 0, 0, 0, 0, time.UTC),
		recipientName: "Vukasin",
	})

	e := email{
		isSubscribed: true,
		body:         "This is a body",
	}

	test1(e, e)

	fmt.Println(getExpenseReport2(email{
		isSubscribed: true,
		body:         "This is my body for the second test from email struct",
		toAddress:    "test@test.com",
	}))

	fmt.Println(getExpenseReport2(sms{
		isSubscribed:  true,
		body:          "This is my body for the second test from sms struct",
		toPhoneNumber: "0625252645",
	}))

	test2(email{
		isSubscribed: true,
		body:         "This is my body for the second test from email struct",
		toAddress:    "test@test.com",
	})

	test2(sms{
		isSubscribed:  true,
		body:          "This is my body for the second test from sms struct",
		toPhoneNumber: "0625252645",
	})

	test2(invalid{})
}
