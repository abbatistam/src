package main

import "fmt"

func main() {
	// Declaracion de variables
	helloMesage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMesage, worldMessage)
	fmt.Println(helloMesage, worldMessage)

	// Printf
	nombre := "Platzi"
	curso := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, curso)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, curso)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, curso)
	fmt.Println(message)

	// Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMesage)
	fmt.Printf("cursos: %T\n", curso)
}
