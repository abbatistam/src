package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// AND

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	// OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	// Convertir texto a numero

	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
}
