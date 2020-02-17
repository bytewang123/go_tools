package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	playMsg := make(chan string)
	go play(playMsg)
	timer := time.NewTimer(3 * time.Second)
	for {
		select {
		case msg := <-playMsg:
			fmt.Println("msg:", msg)
		case <-timer.C:
			fmt.Println("timeout")
		default:
			fmt.Println("default")
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
}

func play(playMsg chan string) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	playMsg <- "playing"
}
