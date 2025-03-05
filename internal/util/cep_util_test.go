package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var successValues = []struct {
	value    string
	expected string
}{
	{value: "12345678", expected: "12345678"},
	{value: "1234567", expected: "01234567"},
	{value: "123456", expected: "00123456"},
	{value: "12345", expected: "00012345"},
	{value: "1234", expected: "00001234"},
	{value: "123", expected: "00000123"},
	{value: "12", expected: "00000012"},
	{value: "1", expected: "00000001"},
	{value: "0", expected: "00000000"},
	{value: "123456789", expected: "123456789"},
}

func TestGetCepArgumentWithSuccessfull(t *testing.T) {
	for _, input := range successValues {
		args := []string{"", input.value}
		cep, err := GetCepArgument(args)
		assert.Nil(t, err)
		assert.Equal(t, input.expected, cep)
	}

}

func TestGetCepArgumentWithInvalidValue(t *testing.T) {
	args := []string{"invalid"}
	_, err := GetCepArgument(args)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "informe o cep")
}
