package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido! Resultado: %d. Valor esperado: %d", total, 30)
	}
}
