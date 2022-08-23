package main

import "testing"

func testeSoma(t *testing.T) {

	total := Soma(10, 100)
	if total != 20 {
		t.Errorf("Resultado da Soma é invalido.")
	}
}
