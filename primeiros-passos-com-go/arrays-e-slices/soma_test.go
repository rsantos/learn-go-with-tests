package arraysslices

import "testing"

func TestSoma(t *testing.T) {
	numeros := []int{1, 2, 3}

	resultado := Soma(numeros)
	esperado := 6

	if esperado != resultado {
		t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
	}
}
