package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Yo yo yo")
	if host, err := os.Hostname(); err == nil {
		fmt.Printf("Hostname: %s\n", host)
	} else {
		fmt.Printf("Error reading hostname %v", err)
	}

	oneSecTicker := time.NewTicker(1 * time.Second)
	for {
		select {
		case currTime := <-oneSecTicker.C:
			fmt.Println(currTime)
		}
	}
}
