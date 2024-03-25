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

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]

		// we cannot assign a value to slice at some index, if that index is of the length of the slice.
		// for example - if we have a slice that is of length 3 (testSlice : = []float64 {1.0, 2.0, 3.0}), we cannot do testSlice[3] = 4.0, that would be an error
		// we can append a new item to the slice (in this case, we are just adding a zero), that will increase the slice length, and therefore, we can add to that slice index
		// so now, we can do something like testSlice[3] += 4.0, we can also reassign the value like testSlice[3] = 4.0
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}

		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func test3(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("======= END Report ========")
}

// slice of slices

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}

	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func test4(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix....\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("====== END Report =======")
}

// tricky slices

// we should not be doing this ===> someSlice = append(otherSlice, element)
// instead, when we append an item to the slice, we should always reassign to that same slice

// range

// example ===> fruits := []string{"apple", "banana", "grape"}
// for i, fruit := range fruits {
//	fmt.Println(i, fruit) => this will output 0 apple, 1 banana, 2 grape
// }

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func test5(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words: \n", msg)

	for _, badWord := range badWords {
		fmt.Println(" -", badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("===================")
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

	test3([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})

	test4(10, 10)

	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "frick", "John"}
	test5(message, badWords)
}
