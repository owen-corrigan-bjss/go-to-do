package main

import (
	"fmt"
	"sync"
)

func addOddNum(num *int) {
	*num += 3
	fmt.Println(*num)
	fmt.Println("odd")
}

func addEvenNum(num *int) {
	*num += 2
	fmt.Println(*num)
	fmt.Println("even")
}

func addOddNumChans(num chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	v := <-num
	v += 3
	fmt.Println(v)
	fmt.Println("odd")
	num <- v
}

func addEvenNumChans(num chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	v := <-num
	v += 2
	fmt.Println(v)
	fmt.Println("even")
	num <- v
}

func main() {
	// //simulating a race with a loop
	// num := 1
	// for i := 0; i <= 10; i++ {
	// 	go addOddNum(&num)
	// 	go addEvenNum(&num)
	// 	time.NewTimer(4 * time.Second)
	// }

	// fmt.Printf("data after the loop %d\n", num)
	var wg sync.WaitGroup
	numChan := make(chan int, 1)
	numChan <- 0
	defer close(numChan)
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go addOddNumChans(numChan, &wg)
		wg.Add(1)
		go addEvenNumChans(numChan, &wg)
	}
	wg.Wait()
}
