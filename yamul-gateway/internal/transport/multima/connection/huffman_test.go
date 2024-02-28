package connection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHuffManCompress(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "nothing",
			input:    []byte{},
			expected: []byte{0b11010000},
		},
		{
			name:     "real example",
			input:    []byte{0xb9, 0x00, 0x00, 0x00, 0x05},
			expected: []byte{0b10110011, 0b00000000, 0b10100011, 0b01000000}, // Checked value by hand
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := make([]byte, len(testCase.expected))
			outputLen := HuffManCompress(result, testCase.input, len(testCase.input))
			assertions := assert.New(t)
			assertions.Equal(len(testCase.expected), outputLen)
			assertions.EqualValues(testCase.expected, result)
		})
	}
}
