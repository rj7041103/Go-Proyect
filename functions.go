package main

import (
"bufio"
"fmt"
"math"
"os"
"strconv"
"strings"
)

func calculateCuadraticFunction(number float64) float64 {
return (number - 3) * (number - 3)
}

func derivate(number float64) float64 {
return 2 * (number - 3)
}

func readFloat(message string) float64 {

    reader := bufio.NewReader(os.Stdin)

    fmt.Print(message)
    input_buffer, _ := reader.ReadString('\n')       // Se leen los datos hasta el salto de linea 
    input := strings.TrimRight(input_buffer, "\r\n") // se quita el salto de linea generado por el usuario

    number, err := strconv.ParseFloat(input, 64) // Convertir la cadena a un número de punto flotante

    if err != nil {
    	fmt.Println("Error al convertir la cadena:", err)
    }

    return number

}

func readNumber(message string) int {
reader := bufio.NewReader(os.Stdin)

    fmt.Print(message)
    input_buffer, _ := reader.ReadString('\n')       // Se leen los datos hasta el salto de linea
    input := strings.TrimRight(input_buffer, "\r\n") // se quita el salto de linea generado por el usuario

    number, err := strconv.Atoi(input) // Convertir la cadena a un número entero

    if err != nil {
    	fmt.Println("Error al convertir la cadena:", err)
    }

    return number

}

func getMinimunValue(number float64, size float64, maxIter int) float64 {
    for i := 0; i < maxIter; i++ {
        // Calcular el gradiente
        gradiente := 2 * (number - 2)

    	// Actualizar el peso
    	x1 := number - (size*gradiente)

    	// se obtiene el valor absoluto de 
    	if math.Abs(x1-number) < 1e-5 {
    		fmt.Println("Mínimo encontrado:", x1)
    		break
    	}

    	// Repetir para la siguiente iteración
    	number = x1
    }

    return number

}

func resolverEcuacionCuadratica(a, b, c float64)  {
	D := math.Pow(b, 2) - 4*a*c

	switch {
	case D > 0:
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		formattedText := fmt.Sprintf("Las raíces son reales y diferentes: %.2f y %.2f", x1, x2)
		fmt.Println(formattedText)
		return 
	case D == 0:
		x := -b / (2 * a)
		formattedText := fmt.Sprintf("La ecuación tiene una raíz real repetida: %.2f", x)
		fmt.Println(formattedText)
		return 
	default:
		realParte := -b / (2 * a)
		imagParte := math.Sqrt(-D) / (2 * a)
		formattedText := fmt.Sprintf("Las raíces son complejas: %.2f ± %.2fi", realParte, imagParte)
		fmt.Println(formattedText)
		return 
	}
}
