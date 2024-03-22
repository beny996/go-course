package main

import (
	"errors"
	"fmt"
)

// arrays

func getMessageWithRetries() [3]string {
	return [3]string{"click here to sign in", "pretty please click here", "we beg you to sign up"}

}

func send(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("sending: %v", msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("complete failure")
		}
	}
}

// slices

const planFree = "free"
const planPro = "pro"

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()

	if plan == planPro {
		// this syntax allows us to use everything from the array, but return type will be a slice!
		return allMessages[:], nil
	}
	if plan == planFree {
		// this is a slice, this slice will take everything from the array between indexes 0 and 2 (not including the item at index 2, but including the one at index 0 )
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

// make

// when we want to make a slice without previously creating an underlying array, we can use the make function

// mySlice := make([]int, 5, 10)

// here, number 5 is saying that the length of the slice will be 5, but we have a total of 10 "slots", if we want to grow our slice

// but we can omit the length of the underlying array, then, the underlying array will have the length of the slice

// mySlice := make([]int, 5)

// we can also create slice literals

// mySlice := []string{"I", "Love", "Go"}

func getMessageCost(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}

func test(messages []string) {
	costs := getMessageCost(messages)

	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf("- %v\n", messages[i])
	}
	fmt.Println("Costs: ")
	for i := 0; i < len(costs); i++ {
		fmt.Printf("- %v\n", costs[i])
	}
}

//variadic functions

func sum(nums ...float64) float64 {
	total := 0.0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total
}

func test2(nums ...float64) {
	total := sum(nums...)

	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("==== End report ====")
}

func main() {
	send("Bob", 0)
	send("Alice", 1)
	send("Mangalam", 2)
	send("Ozgur", 3)

	test([]string{
		"This is a first message",
		"This is test",
		"Testing",
	})

	test2(1, 2, 3, 4, 5, 6)

}
