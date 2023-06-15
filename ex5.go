package main

import (
	"fmt"
)

//Faça um algoritmo que leia o valor do dólar em reais e um valor em dólares, e converta esse valor para reais.

func main() {
	var doli float32
	var dolf float32
	var real float32

	fmt.Print("Qual o valor do dólar em reais?")
	fmt.Scan(&doli)
	fmt.Print("Quantos dólares serão trocados?")
	fmt.Scan(&dolf)

	real = doli * dolf
	fmt.Println("Total recebido em reais: ", real)
}
