package main

import "fmt"

// Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).
func main() {
	a := 0.0
	p := 0.0
	fmt.Print("Peso:")
	fmt.Scan(&p)
	fmt.Print("Altura:")
	fmt.Scan(&a)
	fmt.Printf("IMC: %.2f", p/(a*a))
}
