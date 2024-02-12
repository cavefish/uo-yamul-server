package handlers

import (
	"fmt"
	"testing"
	"yamul-gateway/internal/transport/multima/commands"
	"yamul-gateway/utils/tests/assertions"
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
			assertions.For(t).Equals(testCase.expected, result)
		})
	}
}
