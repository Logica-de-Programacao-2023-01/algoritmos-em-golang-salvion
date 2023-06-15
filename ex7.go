package main

import "fmt"

// Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.
func main() {
	var n int32
	fmt.Print("Digite aqui um número inteiro para descobrir seu antecessor e seu sucessor: ")
	fmt.Scan(&n)
	fmt.Println("Antecessor: ", n-1)
	fmt.Println("Sucessor: ", n+1)
}
