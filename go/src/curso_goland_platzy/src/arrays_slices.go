package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string

	for i:=len(text) -1; i>=0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slices
	slice := []int {0, 1 ,2 ,3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int {8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// Recorriendo slices
	sliceStr := []string {"hola", "que", "hace"}
	for i, valor := range sliceStr {
		fmt.Println(i, valor)
	}

	isPalindromo("amor a roma")
	isPalindromo("casas")
	isPalindromo("Ama")
}
