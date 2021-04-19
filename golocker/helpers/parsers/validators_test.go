package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SimpleTest struct {
	input    string
	expected error
}

//TODO: Add more tests here
//TODO: Test ValidateStringArray
func TestValidateString(t *testing.T) {

	tests := []SimpleTest{
		{
			input:    "qqq",
			expected: nil,
		},
		{
			input:    "",
			expected: fmt.Errorf("Registration information cannot be empty"),
		},
		{
			input:    "qqqq<(.qq|\nq)*q?q>qqqqq",
			expected: fmt.Errorf("Registration information cannot contain: " + "<(.qq|\nq)*q?q>"),
		},
	}

	for _, test := range tests {
		actual := ValidateString(test.input)
		assert.Equal(t, test.expected, actual)
	}

}
