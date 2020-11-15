package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// Sem sinal (Só positivos) -> uint8 uint16 uint32 uint64
	var b byte = 255 // uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// Int com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int64 é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))


	var i2 rune = 'a' // Repesenta um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))


	// Números reais float32 float64
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))//float64


	// Tipo Boolean 
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá meu nome é João"
	fmt.Println(s1 + "!")
	// Em Go não é permitido delimitar um tipo string com aspas simples
	fmt.Println("O tamanho da string é", len(s1))

	//String com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	João`
	fmt.Println("O tamanho da string é", len(s2))


	// Não existe tipo char em Go
	char := 'a' // Não consigo colocar mais de um caracter Ex: 'ab'
	fmt.Println("O tipo char é", reflect.TypeOf(char))// int32
	fmt.Println(char)// Ele retorna um valor da tabela Unicode


}
