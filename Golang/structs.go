package main

import "fmt"

type CadastroKabum struct {
	Nome       string
	Idade      int
	Endereço   string
	Email      string
	Score      float64
	Cadastrado bool
}

type Produto struct {
	ID         int
	Nome       string
	Quantidade int
	Preço      float64
}

func main() {
	// fmt.Println(Pessoa{Nome: "Júlio César"})
	// fmt.Println(Pessoa{Idade: 25})
	P1 := CadastroKabum{Nome: "Júlio César"}
	P1.Idade = 25
	P1.Endereço = "Rua das Flores, 123 - Camopinas - SP"
	P1.Email = "julio.cesar@gmail.com"
	P1.Score = 8.5
	P1.Cadastrado = false

	fmt.Println(P1)

	item := Produto{ID: 1040, Nome: "RTX5090", Quantidade: 1, Preço: 4999.90}
	fmt.Println(item)
}
