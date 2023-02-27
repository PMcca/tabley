package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"tabley/music"
)

var (
	sanitizePattern = regexp.MustCompile(`[-_\s]+`)   // -, _ and whitespace.
	sharpPattern    = regexp.MustCompile(`(?i)sharp`) // Case-insensitive Sharp.
	flatPattern     = regexp.MustCompile(`(?i)flat`)  // Case-insensitive Flat.
)

// ToScale takes a string representing a scale for a given note and returns the corresponding music.Scale.
func ToScale(s string) (*music.Scale, error) {
	if len(s) < 2 {
		return nil, fmt.Errorf("invalid scale %s given", s)
	}
	scaleNormalised := sanitizeInput(s)

	// Get the index of the string to mark for substrings.
	var substringIndex uint8
	var isFlat bool
	switch {
	case scaleNormalised[1] == '#':
		substringIndex = 2
	case strings.ContainsRune(scaleNormalised[1:], '♭'):
		substringIndex = 4 // ♭ is three bytes big.
		isFlat = true      // Flat root note means key signature is minor.
	default:
		substringIndex = 1
		isFlat = music.IsFlatKeySignature(scaleNormalised) // Natural note; Use scale name to determine if flat.
	}

	noteInput := scaleNormalised[:substringIndex]
	scaleName := scaleNormalised[substringIndex:]

	rootNote, err := music.NoteFromString(noteInput)
	if err != nil {
		return nil, err
	}
	intervalPattern, err := music.IntervalPatternFromString(scaleName)
	if err != nil {
		return nil, err
	}

	scale, err := music.NewScale(scaleName, rootNote, intervalPattern, isFlat)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create scale")
	}

	return scale, nil
}

// sanitizeInput returns uppercase s with all whitespaces, hyphens and underscores removed, and accidentals added.
func sanitizeInput(s string) string {
	sanitized := strings.ToUpper(sanitizePattern.ReplaceAllString(s, ""))
	sanitized = sharpPattern.ReplaceAllString(sanitized, "#")
	return flatPattern.ReplaceAllString(sanitized, "♭")
}
