package processor_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
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

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := processor.NewProcessorFromArg(tc.input)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestTabProcessorConvertTabAlignment(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
		tuning   string
	}{
		"SingleDigitToAccidental": {
			input: `
E|--2-2--|
B|--0-0--|`,
			expected: `
E|--F#-F#--|
B|--B--B---|`,
			tuning: "BE",
		},
		//		"ComplexMix": {
		//			input: `
		//E|--------------------------------------------|
		//B|-------18--------18\16----------------------|
		//G|-----------------------------15/18---18-----|
		//D|-------18--------18\16----------------------|
		//A|-----------------------------16/18---18-----|
		//E|--------------------------------------------|`,
		//			expected: `
		//E|--------------------------------------------|
		//B|--------F--------F-\D#----------------------|
		//G|-----------------------------A#/C#---C#-----|
		//D|-------G#--------G#\F#----------------------|
		//A|-----------------------------C#/D#---D#-----|
		//E|--------------------------------------------|`,
		//		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			inputTab := strings.Split(tc.input[1:], "\n")

			proc := processor.NewTabProcessor()
			actual, err := proc.ConvertTab(inputTab, tc.tuning)
			require.NoError(t, err)

			assert.Equal(t, tc.expected[1:], actual)
		})
	}
}

func TestTabProcessorConvertTabValidationError(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		tabInput    []string
		tuningInput string
		expectedErr error
	}{
		"EmptyTabInput": {
			tabInput:    []string{},
			tuningInput: "EADGBE",
			expectedErr: fmt.Errorf("invalid number of tab rows [0]. need at least 1"),
		},
		"EmptyTuning": {
			tabInput:    []string{"E|--2---|"},
			tuningInput: "",
			expectedErr: fmt.Errorf("invalid tuning [] must have at least 1 note"),
		},
		"NumberOfRowsNotMultipleOfNumberOfStrings": {
			tabInput: []string{
				"E|--2---|",
				"B|--2---|",
				"G|--2---|",
			},
			tuningInput: "BE",
			expectedErr: fmt.Errorf("number of strings 2 not divisible by given number tab rows 3"),
		},
		"LengthOfTabRowsNotIdentical": {
			tabInput: []string{
				"E|--2---|",
				"B|--2--|",
			},
			tuningInput: "BE",
			expectedErr: fmt.Errorf("length of tab row [B|--2--|] is not the same as other rows in this group [[E|--2---| B|--2--|]]"),
		},
		"LengthOfTabRowsNotIdenticalInGroupedRow": {
			tabInput: []string{
				// This is valid
				"E|--2---|",
				"B|--2---|",
				// This is invalid
				"E|--1---|",
				"B|--1--|",
			},
			tuningInput: "BE",
			expectedErr: fmt.Errorf("length of tab row [B|--1--|] is not the same as other rows in this group [[E|--1---| B|--1--|]]"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			proc := processor.NewTabProcessor()
			actual, err := proc.ConvertTab(tc.tabInput, tc.tuningInput)

			require.Empty(t, actual)
			assert.ErrorContains(t, err, tc.expectedErr.Error())
		})
	}
}
