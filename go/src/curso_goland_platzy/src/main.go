package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadraro
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// suma
	result := x + y
	fmt.Println("suma:", result)

	// resta
	result = y - x
	fmt.Println("resta:", result)

	// multiplicación
	result = x * y
	fmt.Println("multiplicacion:", result)

	// división
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// area del rectangulo
	base = 10
	altura = 20
	area = base * altura
	fmt.Println("El area del rectangulo es:", area)

	// area del trapecio
	baseMayor := 10
	baseMenor := 8
	altura = 12
	area = ((baseMayor + baseMenor) * altura) / 2
	fmt.Println("La base del trapecio es:", area)

	// area de un circulo
	radio := 3.0
	areaCirculo := pi * (radio * radio)
	fmt.Println("El area del circulo es", areaCirculo)
}
