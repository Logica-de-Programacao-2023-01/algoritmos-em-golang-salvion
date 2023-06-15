package main

import "fmt"

func main() {
	//pedir salário e calcular aumento de 15%
	var sal float32

	fmt.Print("Qual o salário a adicionar aumento? ")
	fmt.Scan(&sal)

	resultado := sal * 1.15
	fmt.Println("Salário final: ", resultado)
}
