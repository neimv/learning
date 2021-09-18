package main

import "fmt"

type car struct {
	brand string
	year int
}

func main() {
	myCar := car{brand: "ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
