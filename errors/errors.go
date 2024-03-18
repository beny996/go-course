package main

import (
	"errors"
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost1, err := sendSMS(msgToCustomer)

	if err != nil {
		return 0.0, err
	}

	cost2, err := sendSMS(msgToSpouse)

	if err != nil {
		return 0.0, err
	}

	return cost1 + cost2, nil

}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send text over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("The sms that costs $%.2f, to be sent to %v cannot be sent", cost, recipient)
}

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
}

//the way to create custom error type is by doing the following :

type userError struct {
	name string
}

func (ue userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", ue.name)
}

//by implementing error interface in the userError struct, userError struct can now be used as a custom error, like this :

func sendSMS1(msg, userName string) (string, error) {
	if len(userName) < 2 {
		return "", userError{name: userName}
	}
	return msg, nil

}

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", de.dividend)
}

//errors package

//we can create a new error with the usage of the package like this :

var err error = errors.New("Something went wrong")

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}

func test1(x, y float64) {
	defer fmt.Println("==========================")

	fmt.Printf("Dividing %.2f by %.2f ... \n", x, y)

	quotient, err := divide(x, y)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f", quotient)
}

func main() {

	value, err := sendSMSToCouple("This is a test ", "This one is okay")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)

	test(1.4, "0525215511")

	test1(2, 0)
}
