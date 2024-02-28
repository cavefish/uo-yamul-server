package handlers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"yamul-gateway/internal/dtos/commands"
)

func TestConvertClientFeaturesToFlags(t *testing.T) {
	cases := []struct {
		input    commands.ClientFeatures
		expected uint32
	}{
		{
			input:    commands.ClientFeatures{},
			expected: 0,
		},
		{
			input: commands.ClientFeatures{
				SingleCharacterSlot: true,
				ContextMenus:        true,
				EnableAOS:           true,
			},
			expected: 0x2C,
		},
	}
	for idx, testCase := range cases {
		t.Run(fmt.Sprintf("case %d", idx), func(t *testing.T) {
			result := ConvertClientFeaturesToFlags(testCase.input)
			assert.New(t).Equal(testCase.expected, result)
		})
	}
}
