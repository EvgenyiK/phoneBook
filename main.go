package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/EvgenyiK/phoneBook/phonebook"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	data, err := phonebook.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString(5)

	t := phonebook.Userdata{
		Username:    random_username,
		Name:        "Evgen",
		Surname:     "Kartapov",
		Description: "This is me!"}

	id := phonebook.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = phonebook.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	// Попытка повторного удаления!
	err = phonebook.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	id = phonebook.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = phonebook.Userdata{
		Username:    random_username,
		Name:        "Evgen",
		Surname:     "Kartapov",
		Description: "Fuck you =)"}

	err = phonebook.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
}
