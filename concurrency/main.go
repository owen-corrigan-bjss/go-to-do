package main

import (
	"fmt"
	"time"
)

func addOddNum(cha chan int) {
	num := <-cha
	if num%2 == 1 {
		num++
	} else {
		num += 2
	}
	fmt.Println(cha)
	cha <- num
}

func addEvenNum(cha chan int) {
	num := <-cha
	if num%2 != 1 {
		num++
	} else {
		num += 2
	}
	fmt.Println(cha)
	cha <- num
}

// type Data struct {
// 	num int
// }

func main() {
	num := make(chan int)

	for i := 0; i <= 10; i++ {
		if i == 0 {
			num <- 0
		}
		go addOddNum(num)
		go addEvenNum(num)

		// go func() { num <- 1 }()
		// go func() { num <- 2 }()
	}
	// fmt.Printf("data after the loop %d\n", num)

}
