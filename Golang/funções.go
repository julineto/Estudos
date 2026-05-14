package main

import "fmt"

func main() {
	nome("Júlio", "César", "Silva")
	fmt.Println(nome("Júlio", "César", "Silva"))

	idade(25)
	fmt.Println(idade(25))

	residência("Americana", "SP")
	fmt.Println(residência("Americana", "SP"))

	trabalho("Analista Acadêmico", "Cogna Educação")
	fmt.Println(trabalho("Analista Acadêmico", "Cogna Educação"))
}

func nome(primeiroNome string, segundoNome string, terceiroNome string) string {
	return primeiroNome + segundoNome + terceiroNome
}

func idade(idade int) int {
	return idade
}

func residência(cidade string, estado string) string {
	return cidade + estado
}

func trabalho(cargo string, empresa string) string {
	return cargo + empresa
}
