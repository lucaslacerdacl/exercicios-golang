package ex2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	_assert := assert.New(t)

	t.Run("Should return the highest value", func(t *testing.T) {
		var givenNumbers = []int{1, 2, 3}

		result, _ := main(givenNumbers)

		_assert.Equal(3, result)
	})

	t.Run("Should return the lowest value", func(t *testing.T) {
		var givenNumbers = []int{10, 8, 6}

		_, result := main(givenNumbers)

		_assert.Equal(6, result)
	})
}
