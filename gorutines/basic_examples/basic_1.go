package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	userId := 102

	responsesChannel := make(chan string)

	go fetchUserData(userId, responsesChannel)
	go fetchUserRecommendations(userId, responsesChannel)

	go()
	for resp := range responsesChannel {
		fmt.Printf("Response %v\n", resp)
	}

	finishTime := time.Since(now)
	fmt.Printf("Time %v\n", finishTime)
}

func fetchUserData(userId int, responsesChannel chan string) {
	// fetch user data

	time.Sleep(1 * time.Second)
	responsesChannel <- "User data fetched"
}

func fetchUserRecommendations(userId int, responsesChannel chan string) {
	time.Sleep(2 * time.Second)
	responsesChannel <- "User recommendations fetched"
}
