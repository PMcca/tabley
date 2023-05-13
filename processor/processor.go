package processor

import "strings"

// MusicProcessor defines the behaviour for building and printing musical constructs in different formats.
type MusicProcessor interface {
	BuildScale()
	ConvertTab(tab []string, tuning string) (string, error)
}

// NewProcessorFromArg returns a BasicProcessor if isBasic is true, or a TabProcessor by default.
func NewProcessorFromArg(isBasic bool) MusicProcessor {
	if isBasic {
		return BasicProcessor{strings.Builder{}}
	}

	return TabProcessor{strings.Builder{}}
}
