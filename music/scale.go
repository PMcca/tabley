package music

import (
	"fmt"
)

// Scale names.
const (
	MajorScale      = "MAJOR"
	MinorScale      = "MINOR"
	MinorNatural    = "MINORNATURAL"
	MinorMelodic    = "MINORMELODIC"
	MajorPentatonic = "MAJORPENTATONIC"
	MinorPentatonic = "MINORPENTATONIC"
)

// IntervalPattern holds the semitone pattern for a scale.
type IntervalPattern []uint8

var (
	intervalMajor           IntervalPattern = []uint8{2, 2, 1, 2, 2, 2, 1}
	intervalMinorNatural    IntervalPattern = []uint8{2, 1, 2, 2, 1, 2, 2}
	intervalMinorMelodic    IntervalPattern = []uint8{2, 1, 2, 2, 2, 2, 1}
	intervalMajorPentatonic IntervalPattern = []uint8{2, 2, 1, 2, 2}
	intervalMinorPentatonic IntervalPattern = []uint8{1, 2, 2, 1, 2}
)

// Map of scale name -> interval pattern
var nameToPattern = map[string]IntervalPattern{
	MajorScale:      intervalMajor,
	MinorScale:      intervalMinorNatural, // Assume "minor" means minor natural.
	MinorNatural:    intervalMinorNatural,
	MinorMelodic:    intervalMinorMelodic,
	MajorPentatonic: intervalMajorPentatonic,
	MinorPentatonic: intervalMinorPentatonic,
}

// Map of scale name -> printable, pretty name
var nameToPretty = map[string]string{
	MajorScale:      "Major",
	MinorScale:      "Minor",
	MinorNatural:    "Minor Natural",
	MinorMelodic:    "Minor Melodic",
	MajorPentatonic: "Major Pentatonic",
	MinorPentatonic: "Minor Pentatonic",
}

// Scale consists of a list of notes and a flag to determine if the scale uses flats or sharps.
type Scale struct {
	Name  string
	Notes []Note
	Flat  bool
}

// NewScale takes a root note and interval pattern and builds a scale of notes from it.
func NewScale(name string, root Note, intervals IntervalPattern, isFlat bool) (*Scale, error) {
	notes := make([]Note, len(intervals)+1)
	notes[0] = root

	nextNote := root
	for i, interval := range intervals {
		n := nextNote.Add(interval)
		notes[i+1] = n
		nextNote = n
	}

	scaleNamePretty, ok := nameToPretty[name]
	if !ok {
		return nil, fmt.Errorf("no scale found for given name %s", name)
	}

	return &Scale{
		Name:  fmt.Sprintf("%s %s", root.StringAccidental(isFlat), scaleNamePretty),
		Notes: notes,
		Flat:  isFlat,
	}, nil
}

// IntervalPatternFromString takes a scale name s and returns the corresponding interval pattern, or error if not found.
func IntervalPatternFromString(scale string) (IntervalPattern, error) {
	p, ok := nameToPattern[scale]
	if !ok {
		return nil, fmt.Errorf("no interval pattern found for scale %scale", scale)
	}
	return p, nil
}

// IsFlatKeySignature returns true for scales that use flat notes, and false for ones that use sharps.
func IsFlatKeySignature(name string) bool {
	switch name {
	case MinorScale, MinorNatural, MinorPentatonic:
		return true
	default:
		return false
	}
}
