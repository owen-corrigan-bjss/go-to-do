package main

import (
	"fmt"
)

func addOddNum(num *int, done chan bool) {
	*num += 3
	fmt.Println(*num)
	fmt.Println("odd")
	done <- true
}

func addEvenNum(num *int, done chan bool) {
	*num += 2
	fmt.Println(*num)
	fmt.Println("even")
	done <- true
}

func main() {
	//simulating a race with a loop
	done := make(chan bool)
	num := 1
	for i := 0; i <= 10; i++ {
		go addOddNum(&num, done)
		<-done
		go addEvenNum(&num, done)
		<-done
	}

	fmt.Printf("data after the loop %d\n", num)

}
