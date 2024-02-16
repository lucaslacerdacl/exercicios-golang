package ex2

type OperatorType string

const (
	HIGH OperatorType = "HIGH"
	LOW  OperatorType = "LOW"
)

func getValueByOperator(numbers []int, operator OperatorType) int {
	var result = numbers[0]

	for _, v := range numbers {
		if operator == HIGH {
			if v > result {
				result = v
			}

		} else if operator == LOW {
			if v < result {
				result = v
			}
		}

	}

	return result
}

/*
 * Faça um programa que recebe três números do usuário, e identifica o maior através de uma função
 * e o menor número através de outra função.
 */
func main(numbers []int) (int, int) {
	highestValue := getValueByOperator(numbers, HIGH)
	lowestValue := getValueByOperator(numbers, LOW)

	return highestValue, lowestValue
}
