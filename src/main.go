package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	
	// Declaración de variables enteras
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

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta 
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// División
	result = y / x
	fmt.Println("División: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Retos
	// - Rectangulo, trapecio y de un círculo

	// Área de un rectangulo
	fmt.Println("Calculo de área de un Rectangulo...")
	baseRectangulo := 50
	alturaRectangulo := 10
	fmt.Println("base: ", baseRectangulo)
	fmt.Println("altura: ", alturaRectangulo)
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("El área de un rectangulo esta dada por la formula base x altura, aplicando esta formula obtenemos: ")
	fmt.Println("Área del rectangulo: ", areaRectangulo)

	// Área de un trapecio
	fmt.Println("Calculo de área de un Trapecio...")
	baseMayorTrapecio := 50
	baseMenorTrapecio := 30
	alturaTrapecio := 20
	fmt.Println("base mayor: ", baseMayorTrapecio)
	fmt.Println("base menor: ", baseMenorTrapecio)
	fmt.Println("altura: ", alturaTrapecio)
	areaTrapecio := ((baseMayorTrapecio + baseMenorTrapecio) * alturaTrapecio) / 2
	fmt.Println("El área de un trapecio esta dada por la suma de sus bases multiplicada por la altura y esto a su vez dividido entre dos, aplicando esta formula obtenemos: ")
	fmt.Println("Área del Trapecio: ", areaTrapecio)

	// Área de un circulo
	fmt.Println("Calculo de área de un Círculo...")
	const PI float64  = 3.14
	radioCirculo := 15
	fmt.Println("radio :", radioCirculo)
	areaCirculo := PI * float64(radioCirculo) * float64(radioCirculo)
	fmt.Println("El área de un circulo esta dada por la multiplicacion de la constante PI por el radio del circulo al cuadrado, aplicando esta formula obtenemos: ")
	fmt.Println("Área del circulo: ", areaCirculo)
}