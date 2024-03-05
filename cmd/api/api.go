package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Testando...")
		time.Sleep(10 * time.Second)
	}
}
