package main

import "fmt"

func main() {
	fmt.Println("operators")

	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 45)
	d := (42 != 43)
	e := (42 < 46)
	f := (42 > 43)

	fmt.Println(a, b, c, d, e, f)
}
