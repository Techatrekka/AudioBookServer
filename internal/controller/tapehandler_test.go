package controller

import (
	"clos/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockTapes() []model.Tape {
	var tapes []model.Tape
	return tapes
}

func TestReturnCatalog(t *testing.T) {
	type Testcase struct {
		//input value
		value string
		//expected values
		expected []byte
	}

	t.Run("function", func(t *testing.T) {

		tests := []Testcase{
			{value: "", expected: []byte("")},
			{value: "", expected: []byte("test finished")},
		}

		for _, test := range tests {
			actual := ReturnCatalog(test.value, mockTapes())
			assert.Equal(t, test.expected, actual)
		}

	})
}
