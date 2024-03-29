package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SimpleTest struct {
	input    string
	expected bool
}

//TODO: Add more tests here
//TODO: Test ValidateStringArray
func TestValidateString(t *testing.T) {
	tests := []SimpleTest{
		{
			input:    "qqq",
			expected: true,
		},
		{
			input:    "",
			expected: false,
		},
		{
			input:    "qqqq<(.qq|\nq)*q?q>qqqqq",
			expected: false,
		},
	}

	for _, test := range tests {
		actual := StringIsValid(test.input)
		assert.Equal(t, test.expected, actual)
	}
}

func TestSanitizeString(t *testing.T) {
	str := fmt.Sprint(`>qqqq <q qqq. q?q q\q q*q q&q q)q q(q q"q q'q q;q q:q q`, "`", "qqqq")

	stripped := SanitizeString(str)

	assert.Equal(t, "qqqq q qqq. qq q\\q qq qq qq qq qq q'q qq qq qqqqq", stripped)

	str = "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq"

	stripped = SanitizeString(str)

	assert.NotEqual(t, str, stripped)
}
