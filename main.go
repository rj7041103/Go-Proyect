package main

import (
    "fmt"
)

// objectiveFunction calcula el valor de la función cuadrática en un punto x.
// En este caso, la función es f(x) = x^2.
func objectiveFunction(x float64) float64 {
    
    return x * x

}

// derivative calcula la derivada de la función cuadrática en un punto x.
// La derivada de f(x) = x^2 es f'(x) = 2x.
func derivative(x float64) float64 {
    return 2 * x
}

// gradientDescent implementa el algoritmo de descenso de gradiente para minimizar la función.
// x0 es el punto inicial, alpha es el tamaño de paso, y maxIter es el número máximo de iteraciones.
// Devuelve el valor de x que minimiza la función y el valor mínimo de la función en ese punto.
func gradientDescent(x0, alpha float64, maxIter int) (float64, float64) {
    // Inicializamos x con el valor inicial x0
    x := x0
    
    // Repetimos el proceso maxIter veces
    for i := 0; i < maxIter; i++ {
        // Calculamos la derivada de la función en el punto actual x
        grad := derivative(x)
        
        // Actualizamos el valor de x moviéndolo en la dirección opuesta a la derivada
        x = x - alpha * grad
    }
    
    // Devolvemos el valor de x que minimiza la función y el valor de la función en ese punto
    return x, objectiveFunction(x)
}

func main() {
    // Punto inicial
    x0 := 10.0
    
    // Tamaño de paso
    alpha := 0.1
    
    // Número máximo de iteraciones
    maxIter := 100
    
    // Llamamos a la función gradientDescent para encontrar el mínimo
    minX, minVal := gradientDescent(x0, alpha, maxIter)
    
    // Imprimimos los resultados
    fmt.Printf("El valor de x que minimiza la función es: %f\n", minX)
    fmt.Printf("El valor mínimo de la función en ese punto es: %f\n", minVal)
}
