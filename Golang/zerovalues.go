package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	i = 25
	fmt.Printf("Inteiro: %v\n", i)

	f = 25.1
	fmt.Printf("Float: %v\n", f)

	b = true
	fmt.Printf("Booleano: %v\n", b)

	s = "Júlio"
	fmt.Printf("String: %q\n", s)

	//Para que o Go não atribua value zero, basta definir o valor da variavel. Assim, retornará o valor definido e não valor zero ou false ou string vazia.
}
