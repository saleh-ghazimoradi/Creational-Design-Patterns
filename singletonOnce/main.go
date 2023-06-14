package main

import "fmt"

func main() {
	for i := 1; i <= 80; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
