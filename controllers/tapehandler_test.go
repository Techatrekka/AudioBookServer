package controllers

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockTapes() []models.Tape {
	var tapes []models.Tape
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
