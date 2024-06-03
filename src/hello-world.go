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

	result := x + x
	//Suma
	fmt.Println("Resultado: ", result)
	//Resta
	fmt.Println("Resta: ", y-z)
	//Multiplicacion
	fmt.Println("Multiplicacion: ", x*int(y))
	//Division
	fmt.Println("Division: ", x/int(y))
	//Modulo
	modulo := x % int(y)
	fmt.Println("Modulo: ", modulo)
	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Declaracion de Variables
	var radio float64 = 14
	//Calculo de Area de un Circulo
	fmt.Print("El area del circulo es: ", pi*radio*radio, "\n")
	//claculo de Area de un Trapecio
	var baseMayor float64 = 10
	var baseMenor float64 = 5
	var alturaTrapecio float64 = 7
	fmt.Print("El area del trapecio es: ", (baseMayor+baseMenor)*alturaTrapecio/2, "\n")
}
