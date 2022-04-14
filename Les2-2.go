package main

import "fmt"

func Sqrt(x float64) float64 {
var a, c float64
	c=3.14
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&a)

	x=4*a/c
	z := 1.0
	check := 0.0
	for i:=0; i <= 100; i++	{
		z = z - (z*z - x)/(2*z)
		if i != 0 {
			if check - z < 0.00000001 {
				break
			}
		}
		check = z
	}
	return z
}

func main() {
	fmt.Println("Диаметр:", Sqrt(5645))
	fmt.Print("Длина окружности: ", Sqrt(5645)*3.14)
}