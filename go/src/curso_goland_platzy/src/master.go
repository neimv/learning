package main

import (
	"fmt"
	pk "mypackage/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	// var myOtherCar pk.carPrivate
	// fmt.Println(myOtherCar)

	pk.PrintMessage("Hello")
	// pk.PrintMessage1("Hello")
}
