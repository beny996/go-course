package main

import (
	"errors"
	"fmt"
	"time"
)

func getLast[T any](slice []T) T {
	if len(slice) == 0 {
		var zeroVal T
		return zeroVal
	}

	return slice[len(slice)-1]
}

type email struct {
	message        string
	senderEmail    string
	recipientEmail string
}

type payment struct {
	amount int
}

func test[T any](slice []T) {
	last := getLast(slice)

	fmt.Println("The last item in the array is: ", last)
}

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	cost := newItem.GetCost()

	if balance < cost {
		return nil, 0.0, errors.New("insufficient funds")
	}

	oldItems = append(oldItems, newItem)

	return oldItems, balance - cost, nil
}

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}

	if s.interval == "yearly" {
		return 250.00
	}

	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail         string
	numEmailIsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailIsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.3
	return float64(otup.numEmailIsAllowed) * costPerEmail
}

func test2[T lineItem](newItem T, oldItems []T, balance float64) {
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("New items are %v\n", newItems)
	fmt.Printf("New balance is %v\n", newBalance)
	fmt.Println("====================")
}

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type userBiller struct {
	PLan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.PLan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.PLan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.
	}

	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func main() {
	test([]email{
		{
			message:        "Test message",
			senderEmail:    "Test sender",
			recipientEmail: "Test recipient",
		},

		{
			message:        "Test message1",
			senderEmail:    "Test sender1",
			recipientEmail: "Test recipient1",
		}})

	test([]payment{
		{
			amount: 123,
		}, {
			amount: 122323,
		}, {
			amount: 123777,
		},
	})

	newItem := oneTimeUsagePlan{
		userEmail:         "test@test.com",
		numEmailIsAllowed: 15,
	}

	oldItems := []oneTimeUsagePlan{
		{
			userEmail:         "test1@test.com",
			numEmailIsAllowed: 13,
		},
		{
			userEmail:         "test2@test.com",
			numEmailIsAllowed: 10,
		},
	}

	newItem1 := subscription{
		userEmail: "newtest@test.com",
		startDate: time.Now(),
		interval:  "monthly",
	}

	oldItems1 := []subscription{
		{
			userEmail: "oldtest@test.com",
			startDate: time.Now(),
			interval:  "yearly",
		},
		{
			userEmail: "oldtest1@test.com",
			startDate: time.Now(),
			interval:  "monthly",
		},
	}

	test2(newItem, oldItems, 120)
	test2(newItem1, oldItems1, 100)
}
