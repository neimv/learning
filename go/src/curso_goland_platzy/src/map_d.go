package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
