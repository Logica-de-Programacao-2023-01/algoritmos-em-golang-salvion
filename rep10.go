package main

import "fmt"

// Faça um algoritmo que leia vários números inteiros e mostre o maior deles.
// A leitura deve ser interrompida quando for lido o valor zero.
func main() {
	var n, a int
	b := 0

	fmt.Print("Quantos números irão ser lidos? ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Número ", i+1, ": ")
		fmt.Scan(&a)
		if a > b {
			b = a
		}
	}
	fmt.Print(b)
}
