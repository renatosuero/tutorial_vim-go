package main

import "testing"

func TestSoma(t *testing.T) {
	n := Soma(600, 66)
	if n != 666 {
		t.Error("Resultado diferente do esperado")
	}
}
func TestSubtracao(t *testing.T) {
	n := Subtracao(10, 5)
	if n != 5 {
		t.Error("Resultado diferente do esperado")
	}
}
func TestMultiplicacao(t *testing.T) {
	n := Multiplicacao(10, 2)
	if n != 20 {
		t.Error("Resultado diferente do esperado")
	}
}
