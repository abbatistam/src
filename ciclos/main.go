package main

import "fmt"

func main() {
	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
