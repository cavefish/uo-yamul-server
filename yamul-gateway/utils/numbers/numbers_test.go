package numbers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigEndianInt(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{
			input:  0x12345678,
			output: 0x78563412,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%x->%x", testCase.input, testCase.output), func(t *testing.T) {
			result := BigEndianInt(testCase.input)
			assert.New(t).Equal(testCase.output, result)
		})
	}
}

func TestBigEndianUInt32(t *testing.T) {
	testCases := []struct {
		input  uint32
		output uint32
	}{
		{
			input:  0x12345678,
			output: 0x78563412,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%x->%x", testCase.input, testCase.output), func(t *testing.T) {
			result := BigEndianUInt32(testCase.input)
			assert.New(t).Equal(testCase.output, result)
		})
	}
}
