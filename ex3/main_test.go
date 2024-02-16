package ex3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	_assert := assert.New(t)

	t.Run("Should return POSITIVE number", func(t *testing.T) {
		result := main(5)

		_assert.Equal(POSITIVE, result)
	})

	t.Run("Should return NEGATIVE number", func(t *testing.T) {
		result := main(-10)

		_assert.Equal(NEGATIVE, result)
	})

	t.Run("Should return ZERO number", func(t *testing.T) {
		result := main(0)

		_assert.Equal(ZERO, result)
	})
}
