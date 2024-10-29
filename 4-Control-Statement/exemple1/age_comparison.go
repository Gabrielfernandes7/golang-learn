package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(int64(time.Now().UTC().UnixNano()))
	var ageJohn, ageKelvim int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John tem ", ageJohn)
	fmt.Println("Kelvim tem", ageKelvim)

	if (ageJohn > ageKelvim) {
		fmt.Println("John é mais velho que Kelvim")
	} else {
		fmt.Println("Kelvim é mais velho que John")
	}
}