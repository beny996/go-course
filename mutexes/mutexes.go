package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	return sc.counts[key]
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

// func (sc safeCounter) slowVal(key string) int {
// 	time.Sleep(time.Microsecond)
// 	return sc.counts[key]
// }

func test(sc safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()

			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)

	}
	sort.Strings(emailsSorted)

	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails \n", email, sc.val(email))
	}
	fmt.Println("========================")
}

type emailTest struct {
	email string
	count int
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mux:    &sync.Mutex{},
	}

	test(sc, []emailTest{
		{
			email: "test@test.com",
			count: 23,
		},
		{
			email: "test@test.com",
			count: 22,
		},
		{
			email: "test2@test.com",
			count: 25,
		},
	})

}
