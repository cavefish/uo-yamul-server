package assertions

import (
	"fmt"
	"testing"
)

func TestAssertionsFor_Equals(t *testing.T) {
	testCases := []struct {
		expected any
		actual   any
		fails    bool
	}{
		{
			expected: 1,
			actual:   1,
			fails:    false,
		},
		{
			expected: 1,
			actual:   2,
			fails:    true,
		},
		{
			expected: int(1),
			actual:   int32(1),
			fails:    true,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("(%v, %v)", testCase.expected, testCase.actual), func(t *testing.T) {
			mock := &testing.T{}
			For(mock).Equals(testCase.expected, testCase.actual)
			if mock.Failed() != testCase.fails {
				t.Fail()
			}
		})
	}
}
