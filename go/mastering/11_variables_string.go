package main

import "fmt"

func main() {
	fmt.Println("string")

	a := `Esto es un 
	string
	literal
	no interpretado
	"¿Lo ves?"`

	fmt.Println(a)
}
