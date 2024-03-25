package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
)

// maps

// ages := make(map[string]int)
// ages["Vukasin"] = 27

// ages := map[string]int {
//	"Vukasin": 27
// }

// when len function is used on maps, it returns the total number of key/value pairs

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	userMap := map[string]user{}

	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}

	return userMap, nil
}

func test(names []string, phoneNumbers []int) {
	userMap, err := getUserMap(names, phoneNumbers)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Creating map...")

	for _, user := range userMap {
		fmt.Printf("key: %s, value: \n", user.name)
		fmt.Printf("- name: %s\n", user.name)
		fmt.Printf("- number: %v\n", user.phoneNumber)
	}
}

type user struct {
	name        string
	phoneNumber int
}

// mutations

// note on passing maps:
// like slices, maps are passed by reference into functions
// this means that when a map is passed into a function we write, we can make changes to the original, we don't have to make a copy
// so, if we modify a map inside of a function, the original map is getting modified

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	user, ok := users[name]

	if !ok {
		return false, errors.New("not found")
	}

	if !user.scheduledForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil

}

func test2(users map[string]user2, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("==========================================")
	deleted, err := deleteIfNecessary(users, name)

	if err != nil {
		fmt.Println(err)
		return
	}

	if deleted {
		fmt.Println("Deleted", name)
		return
	}

	fmt.Println("Did not delete", name)
}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func getCounts(userIDs []string) map[string]int {
	counts := map[string]int{}

	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}

	return counts
}

func test3(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)

	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}

	fmt.Println("===========================")
}

// nested maps

func getNameCounts(names []string) map[rune]map[string]int {
	counts := map[rune]map[string]int{}

	for _, name := range names {
		if name == "" {
			continue
		}

		firstChar := name[0]
		_, ok := counts[rune(firstChar)]
		if !ok {
			counts[rune(firstChar)] = map[string]int{}
		}
		counts[rune(firstChar)][name]++
	}
	return counts
}

func test4(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))

	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("========================")
}

func main() {

	test([]string{"John", "Jack", "Bob"}, []int{123, 345, 678})
	test([]string{"John", "Jack", "Bob"}, []int{123, 345, 678, 901})

	users := map[string]user2{
		"john": {
			name:                 "john",
			number:               123,
			scheduledForDeletion: true,
		},
		"jack": {
			name:                 "jack",
			number:               123446,
			scheduledForDeletion: false,
		},
		"bob": {
			name:                 "bob",
			number:               1231111,
			scheduledForDeletion: true,
		},
	}

	test2(users, "john")
	test2(users, "bob")
	test2(users, "Vukasin")
	test2(users, "jack")

	userIDs := []string{}

	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	test3(userIDs, []string{"00", "ff", "dd"})
	test3(userIDs, []string{"aa", "12", "32"})
	test3(userIDs, []string{"bb", "33"})

}
