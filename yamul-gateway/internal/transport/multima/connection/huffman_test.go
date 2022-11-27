package connection

import (
	"fmt"
	"testing"
	"yamul-gateway/utils/tests/assertions"
)

func TestHuffManCompress(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected []byte
	}{
		{
			input:    []byte{},
			expected: []byte{0x60},
		},
		{
			input:    []byte{0xb9, 0x00, 0x00, 0x00, 0x05},
			expected: []byte{0x59, 0x80, 0x51, 0x80, 0x60},
		},
	}
	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("case_%d", idx), func(t *testing.T) {
			result := make([]byte, len(testCase.expected))
			outputLen := HuffManCompress(result, testCase.input, len(testCase.input))
			assert := assertions.For(t)
			assert.Equals(len(testCase.expected), outputLen)
			assert.EqualList(testCase.expected, result)
		})
	}
}
