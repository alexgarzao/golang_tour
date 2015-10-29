package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		d := ((math.Pow(z, 2) - x) / (2 * z))
		if math.Abs(d) < 0.00000000001 {
			break
		}
		z = z - d
	}
	
	return z
}

func main() {
	values := []float64{2, 3, 5, 9, 25}
	
	for i := 0; i < len(values); i++ {
		value := values[i]
		calculated := Sqrt(values[i])
		expected := math.Sqrt(values[i])
		delta := math.Abs(calculated - expected)
		fmt.Println(
			"Valor=", value,
			"Calculado=", calculated,
			"Esperado=", expected,
			"Delta=", delta,
		)
	}
}
