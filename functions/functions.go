package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}

func getNames() (string, string) {
	return "John", "Doe"
}

func nakedReturn() (x, y, z int) {

	//since we didn't provide values for x, y and z, those will be initialized with the value of 0

	return //this will return x, y and z, since we did not specify the thing we want to return (not recommended to be used)
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilCarRental, yearsUntilDrinking int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	return yearsUntilAdult, yearsUntilCarRental, yearsUntilDrinking

}

func main() {

	fmt.Println(add(10, 3))

	sendsSoFar := 430
	const sendsToAdd = 25

	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)

	fmt.Println("you've sent", sendsSoFar, "messages")

	firstName, _ := getNames()

	fmt.Println("Welcome to Textio", firstName)

	fmt.Println(nakedReturn())

	fmt.Println(yearsUntilEvents(24))

}
