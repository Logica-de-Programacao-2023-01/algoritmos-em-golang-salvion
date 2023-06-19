package main

import "fmt"

//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo.

func main() {
	n := 0
	fmt.Print("Número: ")
	fmt.Scan(&n)
	fmt.Printf("dobro: %d\ntriplo: %d\nquadruplo: %d", n*2, n*3, n*4)
}
