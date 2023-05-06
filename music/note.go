package music

import (
	"fmt"
)

// Note represents a single musical note.
type Note uint8

// Ordered enumeration of notes. Sharps and Flats share the same placement.
const (
	Unknown Note = iota
	A
	As_Bf
	B
	C
	Cs_Df
	D
	Ds_Ef
	E
	F
	Fs_Gf
	G
	Gs_Af

	AStr     string = "A"
	As_BfStr        = "A#/B♭"
	BStr            = "B"
	CStr            = "C"
	Cs_DfStr        = "C#/D♭"
	DStr            = "D"
	Ds_EfStr        = "D#/E♭"
	EStr            = "E"
	FStr            = "F"
	Fs_GfStr        = "F#/G♭"
	GStr            = "G"
	Gs_AfStr        = "G#/A♭"
)

var stringToNote = map[string]Note{
	"A":  A,
	"A#": As_Bf,
	"B♭": As_Bf,
}

// Map of note -> string
var noteToString = map[Note]string{
	A:     AStr,
	As_Bf: As_BfStr,
	B:     BStr,
	C:     CStr,
	Cs_Df: Cs_DfStr,
	D:     DStr,
	Ds_Ef: Ds_EfStr,
	E:     EStr,
	F:     FStr,
	Fs_Gf: Fs_GfStr,
	G:     GStr,
	Gs_Af: Gs_AfStr,
}

// NoteFromString returns the corresponding note for a given string s, or an error if not found.
func NoteFromString(s string) (Note, error) {
	n, ok := stringToNote[s]
	if !ok {
		return Unknown, fmt.Errorf("no note found for input %s", s)
	}
	return n, nil
}

// Add adds a given interval to a note. If the result is 0, 12 (G#/A♭) is returned.
func (n Note) Add(interval uint8) Note {
	sum := Note((uint8(n) + interval) % 12)
	if sum == 0 {
		return Gs_Af // (11+1) % 12 = 0, which should be 12.
	}

	return sum
}

func (n Note) String() string {
	s, ok := noteToString[n]
	if !ok {
		return "unknown"
	}
	return s
}

// StringAccidental returns the given note's accidental, if it exists, determined by isFlat, or the input string otherwise.
func (n Note) StringAccidental(isFlat bool) string {
	s := n.String()
	if len(s) > 1 {
		if isFlat {
			return s[3:]
		}
		return s[:2]
	}

	return s
}
