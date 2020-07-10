package arraysslices

// Soma recebe um array de numeros e retorna a soma deles
func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
