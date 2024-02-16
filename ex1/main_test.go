package ex1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	_assert := assert.New(t)

	t.Run("Should give the sum of given numbers", func(t *testing.T) {
		var givenNumbers = []int{1, 2, 3}

		var expectedResult = givenNumbers[0] + givenNumbers[1] + givenNumbers[2]

		sum, _ := main(givenNumbers)

		_assert.Equal(expectedResult, sum)
	})

	t.Run("Should give the media of given numbers", func(t *testing.T) {
		var givenNumbers = []int{1, 2, 3}

		var expectedResult = (givenNumbers[0] + givenNumbers[1] + givenNumbers[2]) / len(givenNumbers)

		_, media := main(givenNumbers)

		_assert.Equal(expectedResult, media)
	})
}
