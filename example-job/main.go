package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	signal.Ignore()
	fmt.Println("starting")

	duration, err := time.ParseDuration(os.Args[1])
	if err != nil {
		panic(err)
	}
	time.Sleep(duration)
	fmt.Println("completed")
}
