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

func NewTabProcessor() TabProcessor {
	return TabProcessor{
		builder: strings.Builder{},
	}
}

// ConvertTab converts all fret numbers to their corresponding notes, returning the same tab as a string.
func (t TabProcessor) ConvertTab(tabInput []string, tuningInput string) (string, error) {
	if err := validateInput(tabInput, tuningInput); err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}

	tuning, err := music.NotesFromString(tuningInput)
	if err != nil {
		return "", fmt.Errorf("failed to parse tuning from input %s: %w", tuningInput, err)
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

//func convertFretColumn(rows []music.Row, index int) {
//	convertedColumn := make([]string, 0, len(rows))
//	for i := range rows {
//		if unicode.IsDigit(rune(rows[i].Data[index])) {
//
//		}
//	}
//}

// convertFretsFromRow takes a single tab row and converts each fret number to the corresponding note.
func convertFretsFromRow(row string, tuning music.Note) (string, error) {
	b := strings.Builder{}
	rowBytes := []byte(row) // Bytes to catch case of single \ being escaped.

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

	if numOfTabRows == 0 {
		return fmt.Errorf("invalid number of tab rows [%d]. need at least 1", len(tab))
	}

	if numOfStrings == 0 {
		return fmt.Errorf("invalid tuning [%s] must have at least 1 note", tuning)
	}

	// Allow for multiple "sets" of tabs, grouped by numOfStrings. e.g. tuning if == EADGBE, then allow for 12 tab rows,
	// representing 2 sets of 6 strings.
	if numOfTabRows%numOfStrings != 0 {
		return fmt.Errorf("number of strings %d not divisible by given number tab rows %d",
			numOfStrings,
			numOfTabRows)
	}

	lowerBound := 0
	upperBound := numOfStrings
	var tabSet []string
	for upperBound <= len(tab) {
		tabSet = tab[lowerBound:upperBound]
		expectedLength := len(tabSet[0])

		for i := 1; i < len(tabSet); i++ {
			if len(tabSet[i]) != expectedLength {
				return fmt.Errorf("length of tab row [%s] is not the same as other rows in this group [%s]", tabSet[i], tabSet)
			}
		}

		lowerBound = upperBound
		upperBound += numOfStrings
	}
	// TODO check row lengths are the same
	return nil
}
