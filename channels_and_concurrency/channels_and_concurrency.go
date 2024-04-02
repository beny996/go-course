package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {

	// when using go keyword, we are spawning a new go routine. It will execute the code in parallel with the rest of the code (async)

	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)

	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()

	isOld := <-isOldChan

	fmt.Println("email 1 is old:", isOld)

	isOld = <-isOldChan

	fmt.Println("email 2 is old:", isOld)

	isOld = <-isOldChan

	fmt.Println("email 3 is old:", isOld)
}

type email struct {
	date time.Time
}

func test(message string) {
	sendEmail(message)
}

func main() {
	test("Sent")

	filterOldEmails([]email{
		{
			date: time.Now(),
		},
		{
			date: time.Date(2021, 4, 21, 0, 0, 0, 0, time.UTC),
		},
		{
			date: time.Date(2010, 4, 21, 0, 0, 0, 0, time.UTC),
		},
	})
}
