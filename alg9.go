package main

import "fmt"

// Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.
func main() {
	r := 0
	soma := 0
	for i := 1; i < 4; i++ {
		fmt.Printf("%d: ", i)
		fmt.Scan(&r)
		soma += r
	}
	fmt.Printf("A soma é: %d.", soma)
}
