package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	fmt.Println("Creating user", user)
	Stats["create"] = Stats["create"] + 1
}

func UpdateUser(user string) {
	fmt.Println("Updating user", user)
	Stats["update"] = Stats["update"] + 1
}

func PurgeStats() {
	Stats["create"] = 0
	Stats["update"] = 0
}
