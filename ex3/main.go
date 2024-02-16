package ex3

type NumberType string

const (
	POSITIVE NumberType = "POSITIVE"
	NEGATIVE NumberType = "NEGATIVE"
	ZERO     NumberType = "ZERO"
)

/*
 * Faça um programa, com uma função que necessite de um argumento.
 * A função retorna o valor de caractere ‘P’, se seu argumento for positivo,
 * e ‘N’, se seu argumento for zero ou negativo.
 */
func main(givenNumber int) NumberType {
	if givenNumber > 0 {
		return POSITIVE
	} else if givenNumber < 0 {
		return NEGATIVE
	} else {
		return ZERO
	}

}
