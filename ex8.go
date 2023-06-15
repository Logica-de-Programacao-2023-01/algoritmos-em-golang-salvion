package main

import "fmt"

// Faça um algoritmo que leia o número de dias trabalhados por um funcionário e o valor da sua diária e calcule o seu salário.
func main() {
	var diaria, dias float32

	fmt.Println("Qual o valor da diária do seu funcionário?")
	fmt.Scan(&diaria)
	fmt.Println("Quantos dias ele trabalhou?")
	fmt.Scan(&dias)

	r := (diaria * dias)
	fmt.Print("", r)
}
