package music_test

import (
	"github.com/stretchr/testify/require"
	"tabley/music"
	"testing"
)

func TestNotesFromString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input     string
		expected  []music.Note
		errAssert require.ErrorAssertionFunc
	}{
		"NoNoteFoundForStringReturnsError": {
			input:     "This is an invalid input",
			expected:  nil,
			errAssert: require.Error,
		},
		"ReturnsSliceOfNotesForString": {
			input: "EADGBE",
			expected: []music.Note{
				music.E,
				music.A,
				music.D,
				music.G,
				music.B,
				music.E,
			},
			errAssert: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := music.NotesFromString(tc.input)

			tc.errAssert(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}
