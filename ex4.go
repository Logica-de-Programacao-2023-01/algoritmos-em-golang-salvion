package main

import "fmt"

func main() {
	var x1, x2, x3, x4 float64
	fmt.Print("Numero 1: ")
	fmt.Scan(&x1)
	fmt.Print("Numero 2: ")
	fmt.Scan(&x2)
	fmt.Print("Numero 3: ")
	fmt.Scan(&x3)
	fmt.Print("Numero 4: ")
	fmt.Scan(&x4)

	media := (x1 + x2 + x3 + x4) / 4
	fmt.Println("Média: ", media)
}
