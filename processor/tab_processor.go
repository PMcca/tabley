package processor

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"tabley/music"
	"unicode"
)

type TabProcessor struct {
	builder strings.Builder
}

func (t TabProcessor) BuildScale() {
	//TODO implement me
	panic("implement me")
}

// ConvertTab converts all fret numbers to their corresponding notes, returning the same tab as a string.
func (t TabProcessor) ConvertTab(tabInput []string, tuningInput string) (string, error) {
	if err := validateInput(tabInput, tuningInput); err != nil {
		return "", err
	}

	tuning, err := music.NotesFromString(tuningInput)
	if err != nil {
		return "", errors.Wrapf(err, "failed to parse tuning from input %s", tuningInput)
	}

	tab := music.NewTabFromString(tabInput, tuning)

	//numOfTabRows := len(tabInput)
	numOfStrings := len(tuningInput)
	// Get the number of individual sets of rows. e.g. 12 rows of 6 strings = 2 individual tabs.
	//numOfTabs := numOfTabRows / numOfStrings

	for i, r := range tab.Rows {
		if i == (numOfStrings) {
			t.builder.WriteString("")
			continue
		}

		row, err := convertFretsFromRow(r.Data, r.Tuning)
		if err != nil {
			return "", errors.Wrapf(err, "failed to convert tab row's %s frets to notes", row)
		}

		t.builder.WriteString(fmt.Sprintf("%s\n", row))
	}

	return t.builder.String(), nil
}

// convertFretsFromRow takes a single tab row and converts each fret number to the corresponding note.
func convertFretsFromRow(row string, tuning music.Note) (string, error) {
	b := strings.Builder{}
	rowBytes := []byte(row) // Catches case of single \ being escaped.

	for i := 0; i < len(rowBytes); i++ {
		if unicode.IsDigit(rune(rowBytes[i])) {
			fretNum := ""
			if i < len(rowBytes)-1 && unicode.IsDigit(rune(row[i+1])) { // Check if next number is also a digit.
				fretNum = string(rowBytes[i : i+2])
				i++
			} else {
				fretNum = string(rowBytes[i])
			}

			fretInt, err := strconv.Atoi(fretNum)
			if err != nil {
				return "", errors.Wrapf(err, "failed to convert fretNum %s to int", fretNum)
			}

			convertedNote := music.NoteFromFretNumber(uint8(fretInt), tuning)

			noteStr := convertedNote.String()
			if len(noteStr) < 1 && !unicode.IsDigit(rune(rowBytes[i+1])) {
				noteStr += "-"
			}
			b.WriteString(noteStr)
		} else {
			x := rowBytes[i]
			b.WriteByte(x)
		}
	}

	return b.String(), nil
}

func validateInput(tab []string, tuning string) error {
	numOfTabRows := len(tab)
	numOfStrings := len(tuning)

	if numOfTabRows < 2 {
		return fmt.Errorf("invalid number of tab rows %d. need at least 2", len(tab))
	}

	if numOfStrings < 2 {
		return fmt.Errorf("invalid tuning %s must have at least 2 notes", tuning)
	}

	if numOfTabRows%numOfStrings != 0 {
		return fmt.Errorf("number of strings %d not divisible by given number tab rows %d, check your input",
			numOfStrings,
			numOfTabRows)
	}

	return nil
}
