package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(4, 2, 5)
	esperado := 11
	if resultado != esperado {
		t.Error("Esperado:", esperado, "Obtido:", resultado)

	}
}
func TestMultiplicacao(t *testing.T) {
	resultado := multiplicacao(10, 20)
	esperado := 200
	if resultado != esperado {
		t.Error("Esperado:", esperado, "Obtido:", resultado)
	}
}
func TestSubtracao(t *testing.T) {
	resultado := subtracao(10, 5)
	esperado := 7
	if resultado != esperado {
		t.Error("Esperado:", esperado, "Obtido:", resultado)
	}
}
func TestDivisao(t *testing.T) {
	resultado := divisao(10, 2)
	esperado := 5
	if resultado != esperado {
		t.Error("Esperado:", esperado, "Obtido:", resultado)
	}
}
