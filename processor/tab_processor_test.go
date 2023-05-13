package processor_test

import (
	"github.com/stretchr/testify/require"
	"tabley/processor"
	"testing"
)

func TestTabProcessor_ConvertTab(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		tabInput    []string
		tuningInput string
		expected    string
		errAssert   require.ErrorAssertionFunc
	}{
		"LessThan2TabRowsReturnsError": {
			tabInput:  []string{"E--2----4-3----"},
			expected:  "",
			errAssert: require.Error,
		},
		"LessThan2NotesInTuningReturnsError": {
			tuningInput: "A",
			expected:    "",
			errAssert:   require.Error,
		},
		"MismatchOfNumberOfRowsAndTuningReturnsError": {
			tabInput: []string{
				"1",
				"2",
				"3",
				"4",
				"5",
				"6",
			},
			tuningInput: "EAD",
			expected:    "",
			errAssert:   require.Error,
		},
		"FailureToParseTuningReturnsError": {
			tuningInput: "XYZMAS",
			expected:    "",
			errAssert:   require.Error,
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			proc := processor.NewProcessorFromArg(false)

			actual, err := proc.ConvertTab(tc.tabInput, tc.tuningInput)

			tc.errAssert(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}
