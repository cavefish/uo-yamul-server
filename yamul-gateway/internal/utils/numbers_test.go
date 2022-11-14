package utils

import (
	"fmt"
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
			if result != testCase.output {
				t.Error(fmt.Sprintf("%x is different than expected %x", result, testCase.output))
			}
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
			if result != testCase.output {
				t.Error(fmt.Sprintf("%x is different than expected %x", result, testCase.output))
			}
		})
	}
}
