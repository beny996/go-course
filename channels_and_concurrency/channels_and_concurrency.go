package main

import (
	"fmt"
	"math/rand"
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

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {

		//this will wait until we receive a token from the channel and then go to the next iteration
		<-dbChan
	}
}

func test2(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases... \n", numDBs)
	waitForDbs(numDBs, dbChan)
	fmt.Println("All databases are online")
}

func getDatabasesChannel(numbDBS int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numbDBS; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}

func addEmailsToQueue(emails []string) chan string {
	// this is a buffered channel, second parameter to the make function will set the buffer size for the channel
	emailsToSend := make(chan string, len(emails))

	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email: ", email)

	}
}

func test3(emails ...string) {
	fmt.Println("==============================")
	fmt.Printf("Adding %v emails to queue... \n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==============================")

}

func countReports(numSentCh chan int) int {
	total := 0

	// infinite loop
	for {
		numSent, ok := <-numSentCh
		if !ok {
			break

		}
		total += numSent
	}

	return total
}

func test4(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent! \n", numReports)
	fmt.Println("========================")

}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batches of %v reports \n", numReports)
		time.Sleep(time.Microsecond * 100)
	}
	close(ch)
}

func concurrentFib(n int) {
	chInts := make(chan int)
	go func() {
		fibonacci(n, chInts)
	}()

	for v := range chInts {
		fmt.Println(v)
	}

}

func test5(n int) {
	fmt.Printf("Printing %v numbers... \n", n)
	concurrentFib(n)
	fmt.Println("==============================")
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

// select

// select is like a switch statement but for channels

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("EMAIL: ", email)
}

func test6(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("======================")

}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(time.Millisecond * 500)

		}
	}
}

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved")

}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test7() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===================")
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
	test2(3)

	test3("First email", "Second email", "Third email")

	test4(3)

	test5(8)

	test6([]string{
		"First sms",
		"Second sms",
		"Third sms",
	}, []string{
		"First email",
		"Second email",
		"Third email",
	})

	test7()
}
