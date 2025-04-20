package music_test

import (
	"github.com/stretchr/testify/require"
	"tabley/music"
	"testing"
)

var (
	majorInterval = music.IntervalPattern{2, 2, 1, 2, 2, 2, 1}
)

func TestNewScale(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		scaleName       string
		rootNote        music.Note
		intervalPattern music.IntervalPattern
		isFlat          bool
		expected        *music.Scale
		errAssert       require.ErrorAssertionFunc
	}{
		"UnknownPrettyNameMappingReturnsError": {
			scaleName:       "some fake scale name",
			rootNote:        music.A,
			intervalPattern: music.IntervalPattern{},
			expected:        nil,
			errAssert:       require.Error,
		},
		"BuildsMajorScale": {
			scaleName:       music.MajorScale,
			rootNote:        music.C,
			intervalPattern: majorInterval,
			isFlat:          false,
			expected: &music.Scale{
				Name:  "C Major",
				Notes: []music.Note{music.C, music.D, music.E, music.F, music.G, music.A, music.B, music.C},
				Flat:  false,
			},
			errAssert: require.NoError,
		},
		"BuildsMajorScaleWithSharp": {
			scaleName:       music.MajorScale,
			rootNote:        music.Gs_Af,
			intervalPattern: majorInterval,
			isFlat:          false,
			expected: &music.Scale{
				Name:  "G# Major",
				Notes: []music.Note{music.Gs_Af, music.As_Bf, music.C, music.Cs_Df, music.Ds_Ef, music.F, music.G, music.Gs_Af},
				Flat:  false,
			},
			errAssert: require.NoError,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := music.NewScale(tc.scaleName, tc.rootNote, tc.intervalPattern, tc.isFlat)

			tc.errAssert(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestIntervalPatternFromString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input     string
		expected  music.IntervalPattern
		errAssert require.ErrorAssertionFunc
	}{
		"UnknownScaleNameReturnsError": {
			input:     "some fake scale",
			expected:  nil,
			errAssert: require.Error,
		},
		"MajorScaleIntervalPatternFromString": {
			input:     music.MajorScale,
			expected:  majorInterval,
			errAssert: require.NoError,
		},
		"MinorScaleIntervalPatternFromString": {
			input:     music.MinorScale,
			expected:  music.IntervalPattern{2, 1, 2, 2, 1, 2, 2},
			errAssert: require.NoError,
		},
		"MinorNaturalScaleIntervalPatternFromString": {
			input:     music.MinorNatural,
			expected:  music.IntervalPattern{2, 1, 2, 2, 1, 2, 2},
			errAssert: require.NoError,
		},
		"MinorMelodicScaleIntervalPatternFromString": {
			input:     music.MinorMelodic,
			expected:  music.IntervalPattern{2, 1, 2, 2, 2, 2, 1},
			errAssert: require.NoError,
		},
		"MajorPentatonicScaleIntervalPatternFromString": {
			input:     music.MajorPentatonic,
			expected:  music.IntervalPattern{2, 2, 1, 2, 2},
			errAssert: require.NoError,
		},
		"MinorPentatonicScaleIntervalPatternFromString": {
			input:     music.MinorPentatonic,
			expected:  music.IntervalPattern{1, 2, 2, 1, 2},
			errAssert: require.NoError,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := music.IntervalPatternFromString(tc.input)

			tc.errAssert(t, err)
			require.Equal(t, tc.expected, actual)

		})
	}
}
