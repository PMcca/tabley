package music

// Tab represents a musical tablature, consisting of an ordered collection of rows.
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
	tab := Tab{}
	notesIndex := len(tuning) - 1
	for _, rawRow := range rawTab {
		if notesIndex < 0 {
			notesIndex = len(tuning) - 1
		}

		row := Row{
			Data:   rawRow,
			Tuning: tuning[notesIndex],
		}
		tab.Rows = append(tab.Rows, row)
		notesIndex--
	}

	return tab
}
