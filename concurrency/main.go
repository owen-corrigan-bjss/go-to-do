package main

import (
	"fmt"
	"time"
)

func addOddNum(data *Data) {
	if data.num%2 == 1 {
		data.num++
	} else {
		data.num += 2
	}
	fmt.Println(data.num)
}

func addEvenNum(data *Data) {
	if data.num%2 != 1 {
		data.num++
	} else {
		data.num += 2
	}
	fmt.Println(data.num)
}

type Data struct {
	num int
}

func main() {
	data := Data{1}

	for i := 0; i <= 10; i++ {
		go addOddNum(&data)
		go addEvenNum(&data)
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("data after the loop %d\n", data.num)

}
