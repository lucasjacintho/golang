package main

import (
	"fmt"
	// Criando uma referencia do import
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	//Forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// Forma de declarar bloco de variaveis
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)



// Forma de declarar variaveis em uma unica linha
// e = true 
// f = false
	var e, f bool = true, false
	fmt.Println(e, f)


// g = 2
// h = false
// i = epa!
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
