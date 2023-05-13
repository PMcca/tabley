package music

// Tab represents a musical tablature, consisting of a ordered collection of rows.
type Tab struct {
	Rows []Row
}

// Row represents a single string in a tab, holding the fret numbers and the string's tuning.
type Row struct {
	Data   string
	Tuning Note
}

// NewTabFromString takes an array of strings holding the raw tab and returns a Tab object from it.
// The returned tab will be in reverse order from the given tuning. e.g. tuning = EADGBE will return a tab with rows
// in the order EBGDAE.
func NewTabFromString(rawTab []string, tuning []Note) Tab {
	t := Tab{}
	notesIndex := len(tuning) - 1
	for _, r := range rawTab {
		if notesIndex < 0 {
			notesIndex = len(tuning) - 1
		}

		row := Row{
			Data:   r,
			Tuning: tuning[notesIndex],
		}
		t.Rows = append(t.Rows, row)
		notesIndex--
	}

	return t
}
