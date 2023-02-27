package processor_test

import (
	"github.com/stretchr/testify/require"
	"tabley/processor"
	"testing"
)

func TestNewProcessorFromArg(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    bool
		expected processor.MusicProcessor
	}{
		"TrueArgReturnsBasicProcessor": {
			input:    true,
			expected: processor.BasicProcessor{},
		},
		"FalseArgReturnsBasicProcessor": {
			input:    false,
			expected: processor.TabProcessor{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := processor.NewProcessorFromArg(tc.input)
			require.Equal(t, tc.expected, actual)
		})
	}
}
