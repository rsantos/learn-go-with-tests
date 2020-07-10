package iteracao

// Repetir recebe um caractere e retorna ele repetido 5 vezes
func Repetir(caractere string, numeroRepeticoes int) string {
	var repeticoes string
	for i := 0; i < numeroRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
