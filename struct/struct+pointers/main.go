package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")

}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(*b)

	*b = 100

	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}

	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

}
