package main

import "fmt"

var msg = "Projeto usado para mostrar o incrível plugin vim-go"

func main() {
	fmt.Println(msg)
}

//Função que soma dois números inteiros
func Soma(x, y int) int {
	return x + y
}

//Função que subtrai dois números inteiros
func Subtracao(x, y int) int {
	return x - y
}

//Função que multiplica dois números inteiros
func Multiplicacao(x, y int) int {
	return x * y
}

//Função que divide dois números inteiros
func Divisao(x, y int) int {
	return x / y
}
