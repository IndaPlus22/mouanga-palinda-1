package main

import (
	"fmt"
	"time"
)

func main() {
	go Remind("Time to eat", 10)
	go Remind("Time to work", 30)
	Remind("Time to sleep", 60)
}

func Remind(text string, delay time.Duration) {
	for {
		time.Sleep(delay * time.Second)
		cur_time := time.Now().Format("15:04:05")
		fmt.Printf("The time is %s: %s\n", cur_time, text)
	}
}
