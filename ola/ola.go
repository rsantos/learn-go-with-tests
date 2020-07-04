package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

// Ola function return string Olá, mundo
func Ola(nome string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
