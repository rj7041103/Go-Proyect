package main

import (
	"fmt"
)

func main() {
	// Punto inicial x0, tamaño de paso alpha, número máximo de iteraciones maxIter
	/* x0 := 0.0
	alpha := 0.01
	maxIter := 1000 */

	number := readFloat("Ingresa un valor para x para calcular:  F(x) = (x - 3)^2: ")
	size_step:= readFloat("Ingrese el peso de ajuste del paso: ")
	max_iter := readNumber("Ingrese el maximo de iteraciones: ")

	fmt.Println("El resultado de la función F(x) = (x - 3)^2 es:", calculateCuadraticFunction(number))
	fmt.Println("El resultado de la derivada de la función F'(x) = 2(x - 3) es:", derivate(number))
	resolverEcuacionCuadratica(1,-6,9)
	// Ejecutar el descenso de gradiente y obtener el valor mínimo de x y f(x)
	minimo := getMinimunValue(number, size_step, max_iter)
	fmt.Println("El valor que minimiza la funcion es x =", minimo)
	fmt.Println("El valor minimo de la funcion es F(x) =", calculateCuadraticFunction(minimo))
}
