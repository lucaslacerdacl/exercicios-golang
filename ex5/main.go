package ex5

/*
 * Crie uma função que receba um valor e uma porcentagem como parâmetros. A função deve retornar o valor acrescido da porcentagem indicada.
 */
func main(givenNumber float64, givenPercentage float64) float64 {

	percentageFromNumber := givenNumber * givenPercentage / 100

	return givenNumber + percentageFromNumber
}
