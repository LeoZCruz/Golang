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
	//Valores Primitivos
	var entero8 int8 = 127
	var entero16 int16 = 32767
	var entero32 int32 = 2147483647
	var entero64 int64 = 9223372036854775807
	fmt.Println(entero8, entero16, entero32, entero64)
	//Valores Primitivos
	var entero8Negativo int8 = -128
	var entero16Negativo int16 = -32768
	var entero32Negativo int32 = -2147483648
	var entero64Negativo int64 = -9223372036854775808
	fmt.Println(entero8Negativo, entero16Negativo, entero32Negativo, entero64Negativo)
	//Valores Primitivos
	var enteroSinSigno8 uint8 = 255
	var enteroSinSigno16 uint16 = 65535
	var enteroSinSigno32 uint32 = 4294967295
	var enteroSinSigno64 uint64 = 18446744073709551615
	fmt.Println(enteroSinSigno8, enteroSinSigno16, enteroSinSigno32, enteroSinSigno64)
	//Valores Primitivos
	var flotante32 float32 = 3.40282346638528859811704183484516925440e+38
	var flotante64 float64 = 1.797693134862315708145274237317043567981e+308
	fmt.Println(flotante32, flotante64)
	//Valores Primitivos
	var complejo64 complex64 = 1 + 2i
	var complejo128 complex128 = 1 + 2i
	fmt.Println(complejo64, complejo128)
	//Valores Primitivos
	var cadena string = "Hola Mundo"
	fmt.Println(cadena)
	//Valores Primitivos
	var caracter byte = 'A'
	fmt.Println(caracter)
	//Valores Primitivos
	var bandera bool = true
	fmt.Println(bandera)
	//Valores Primitivos
	var error error = nil
	fmt.Println(error)
	//Valores Primitivos
	var puntero *int = &x
	fmt.Println(puntero)
	//Valores Primitivos
	var arreglo [3]int
	fmt.Println(arreglo)
	//Valores Primitivos
	var slice []int
	fmt.Println(slice)
	//Valores Primitivos
	var mapa map[string]int
	fmt.Println(mapa)
	//Valores Primitivos
	var canal chan int
	fmt.Println(canal)

}
