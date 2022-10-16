package main

import (
	pk "cursoBasicoPlatzi/src/struct/modificadores/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020

	fmt.Println(myCar)

	pk.PrintMessage("Hola platzi")

}
