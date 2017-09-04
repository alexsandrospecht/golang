package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 0

	for {
		n++

		i := rand.Intn(4200)
		fmt.Println(i)

		if i % 42 == 0 {
			break
		}
	}
	fmt.Println("Saída após %d iterações.", n)
}
