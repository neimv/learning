package main

import "fmt"

// Estilo clases
type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func (myPC pc) String() string {
	return fmt.Sprintf("tengo %d GB de RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

// fin de estilo clases

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b) // parecido a C con *

	*b = 100
	fmt.Println(a)

	myPc := pc{
		ram:   16,
		disk:  200,
		brand: "msi",
	}
	fmt.Println(myPc)
	myPc.ping()

	// Update cosas
	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)

	// Personalizando salidas

}
