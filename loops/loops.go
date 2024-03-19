package main

import "fmt"

func bulkSend(numMessages int) float64 {
	cost := 0.0

	for i := 0; i < numMessages; i++ {
		cost += 1 + (0.01 * float64(i))
	}

	return cost

}

func test(numMessages int) {
	fmt.Printf("Sending %v messages \n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = $%.2f\n", cost)
	fmt.Println("===================")
}

//omitting conditions
// we can omit any of the conditions from the for loop
func maxMessages(thresh float64) int {
	cost := 0.0

	for i := 0; ; i++ {
		cost += 1.0 + (0.01 * float64(i))
		if cost > thresh {
			return i
		}

	}
}

func test2(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: %v\n", max)
	fmt.Println("========================")
}

// while loops (there is no explicit while loops in go)
// to create a loop that is like a while loop in other languages, we can just provide a condition to the for loop, like in the example below :

func exampleWhileLoop() {
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Printf("Plant height is %v", plantHeight)
		plantHeight++
	}
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0

	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}
}

func printPrimes(max int) {
	for i := 2; i < max+1; i++ {
		if i == 2 {
			fmt.Println(i)
			continue
		}
		if i%2 == 0 {
			continue
		}
		isPrime := true
		for n := 3; n*n < i+1; n++ {
			if i%n == 0 {
				isPrime = false
				break
			}

		}
		if !isPrime {
			continue
		}
		fmt.Println(i)
	}
}

func test3(max int) {
	fmt.Printf("Primes up to %v\n", max)
	printPrimes(max)
	fmt.Println("=======================")
}

func main() {
	test(10)

	test2(10)

	fizzBuzz()

	test3(10)
	test3(20)

}
