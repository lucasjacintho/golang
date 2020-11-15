package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// Go não permite concatenar um valor String com um Int
	//fmt.Println("O valor de x é " + x)

	//Para isso devo converter o valor para String 
	// e atribuir a uma variavel
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)


	// Forma correta de concatenar uma String com Int 
	// sem precisar converter
	fmt.Println("O valor de x é ", x)

	//Outra forma seria utilizando o template string
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	
	// Tipos de template strings
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	// Forma generica do template strings
	fmt.Printf("\n %v %v %v %v %v", a, b, b, c, d)
}
