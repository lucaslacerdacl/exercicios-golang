package ex4

import "strings"

type TextCaseType string

const (
	UPPER  TextCaseType = "UPPER"
	LOWER  TextCaseType = "LOWER"
	NORMAL TextCaseType = "NORMAL"
)

/*
 * Crie uma função que receba uma string como parâmetro e retorne a mesma string com todas as letras em maiusculo.
 * Faça uma segunda função que retorne em minusculo.
 */
func main(givenText string, operator TextCaseType) string {
	if operator == UPPER {
		return strings.ToUpper(givenText)
	} else if operator == LOWER {
		return strings.ToLower(givenText)
	} else {
		return givenText
	}
}
