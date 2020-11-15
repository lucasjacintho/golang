package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	// Valores zero são valores padrão para quando as variaveis não são inicializadas
	fmt.Printf("%v %v %v   %q  %v", a, b, c, d, e)
	//Output: 	0  0 false "" <nil> 

}
