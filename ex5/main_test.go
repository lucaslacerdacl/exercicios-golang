package ex5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	_assert := assert.New(t)

	t.Run("Should plus 5 percentage from 10", func(t *testing.T) {
		result := main(10, 5)

		_assert.Equal(10.5, result)
	})

	t.Run("Should plus 10 percentage from 320.5", func(t *testing.T) {
		givenNumber := 320.5

		result := main(givenNumber, 10)

		_assert.Equal(givenNumber+32.05, result)
	})

}
