package main

import (
	"fmt"
)

func test(done chan bool) {
	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}
	done <- true
}
func main() {
	done := make(chan bool)
	go test(done)
	<-done
}
