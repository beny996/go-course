package main

import "fmt"

func main() {
	// one way of defining a variable

	// var number uint8 = 16

	// var string1 string = "testString"

	// var boolean bool = true

	// //other way of defining variables is by letting the variable infer the type itself (with the shorthand syntax)

	// number2 := 22

	// // string2 := "Testing"

	// // boolean2 := false

	// fmt.Println(number, string1, boolean)
	// fmt.Printf("The type of the number2 variable is %T\n", number2)

	// accountAge := 2.6

	// fmt.Println("Your account has existed for", accountAge, "years")

	// accountAgeInt := int(accountAge)

	// fmt.Println("Your account has existed for", accountAgeInt, "years")

	const premiumPlanName = "Premium plan"
	const basicPlanName = "Basic plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)

	fmt.Printf("I am %v years old \n", 28)

	const name = "Vukasin"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)

	messageLen := 10
	maxMessageLen := 20

	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

}
