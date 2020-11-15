package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	// Convertendo valores para inteiro
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) // Aqui é feito o arredondamento ignorando as casas decimais ou seja 6.9 -> 6

	// CUIDADO
	// Um valor inteiro no momento do print precisa ser convertido para o
	// tipo string, porém dessa forma abaixo ele vai me trazer o valor na
	// Tabela Unicode representado por esse número
	fmt.Println("Teste " + string(97))

	// Maneira correta de conversão
	fmt.Println("Teste " + strconv.Itoa(97))

	// Conversão de string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	// Algumas funções em Go irão me retornar dois valores
	// primeiro será o meu objeto convertido nesse caso o 123
	// o segundo valor seria de erro caso eu tivesse colocado um valor não inteiro
	// o _ está falando que eu quero ignorar esse segundo valor


	// Conversão de string para bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
