package main

import "fmt"

func main() {
	//Declaracion de Constantes
	const pi float64 = 3.1415
	const pi2 = 3.1415
	//Tipos de datos
	fmt.Println("PI_1:", pi, "\nPI_2:", pi2)
	//Declaracion de Variables Enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)
	fmt.Print("Area: ", base*altura, "\n")
	//Zero Values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	fmt.Println("Hello, World!")

	//Calculo de Area de un Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Print("El area del cuadrado es: ", areaCuadrado, "\n")
	x := 10
	y := 4.44
	z := float64(x) + y
	fmt.Println("Suma: ", z)
}
