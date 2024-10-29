package main

import (
	"fmt"
	"time"
)

func main()  {
	n := 1000000
	inicio := time.Now()

	for i := 0; i < n; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println(time.Now().Sub(inicio))
}