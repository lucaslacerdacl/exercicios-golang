package ex4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	_assert := assert.New(t)

	t.Run("Should return upper text", func(t *testing.T) {
		result := main("hello world", UPPER)

		_assert.Equal("HELLO WORLD", result)
	})

	t.Run("Should return lower text", func(t *testing.T) {
		result := main("HELLO WORLD", LOWER)

		_assert.Equal("hello world", result)
	})

	t.Run("Should return normal text", func(t *testing.T) {
		result := main("Hello World", NORMAL)

		_assert.Equal("Hello World", result)
	})

}
