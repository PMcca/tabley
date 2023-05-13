package music_test

import (
	"github.com/stretchr/testify/require"
	"tabley/music"
	"testing"
)

func TestNewTabFromString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		rawRows  []string
		tuning   []music.Note
		expected music.Tab
	}{
		"BuildsTabFromStringAndTuning": {
			rawRows: []string{
				"--4--12--5-3--",
				"--4----------",
				"--2----------",
				"-------------",
				"------------9",
				"------------8",
			},
			tuning: []music.Note{
				music.E,
				music.A,
				music.D,
				music.G,
				music.B,
				music.E,
			},
			expected: music.Tab{
				Rows: []music.Row{
					{
						Data:   "--4--12--5-3--",
						Tuning: music.E,
					},
					{
						Data:   "--4----------",
						Tuning: music.B,
					},
					{
						Data:   "--2----------",
						Tuning: music.G,
					},
					{
						Data:   "-------------",
						Tuning: music.D,
					},
					{
						Data:   "------------9",
						Tuning: music.A,
					},
					{
						Data:   "------------8",
						Tuning: music.E,
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := music.NewTabFromString(tc.rawRows, tc.tuning)

			require.Equal(t, tc.expected, actual)
		})
	}
}
