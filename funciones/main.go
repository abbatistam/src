package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)

}

func returnValue(a int) int {
	return a * 2
}

func doubleValue(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "hola")
	value := returnValue(2)
	fmt.Println("value:", value)

	value1, value2 := doubleValue(2)
	fmt.Println(value1, value2)
}
