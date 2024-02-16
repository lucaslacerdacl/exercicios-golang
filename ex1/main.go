package ex1

import "fmt"

func sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func media(numbers []int) int {
	var sum = sum(numbers)

	var length = len(numbers)

	return sum / length
}

/*
 * Faça um programa, com uma função que necessite de três argumentos,
 * e que forneça a soma desses três argumentos através de uma função.
 * Seu script também deve fornecer a média dos três números, através de uma segunda função que chama a primeira.
 */
func main(givenNumbers []int) (int, int) {
	sumGivenNumbers := sum(givenNumbers)
	fmt.Println("Soma dos números: $d", sumGivenNumbers)

	mediaGivenNumbers := media(givenNumbers)
	fmt.Println("Soma dos números: $d", mediaGivenNumbers)

	return sumGivenNumbers, mediaGivenNumbers
}
