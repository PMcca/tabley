package processor

// MusicProcessor defines the behaviour for building and printing musical constructs in different formats.
type MusicProcessor interface {
	BuildScale()
}

// NewProcessorFromArg returns a BasicProcessor if arg is true, or a TabProcessor by default.
func NewProcessorFromArg(arg bool) MusicProcessor {
	if arg {
		return BasicProcessor{}
	}

	return TabProcessor{}
}
