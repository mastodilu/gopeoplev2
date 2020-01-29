package mysignals

// MySignal represents a signal
type MySignal int

const (
	// StartLife signal sent when a new person is born
	StartLife MySignal = 1 + iota
)

func (ms MySignal) String() string {
	switch ms {
	case StartLife:
		return "START_LIFE"
	}
	return "unknown signal"
}
