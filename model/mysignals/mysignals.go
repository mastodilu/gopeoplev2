package mysignals

// LifeSignal represents a signal
type LifeSignal int

const (
	// StartLife signal sent when a new person is born
	StartLife LifeSignal = 1 + iota
)

func (ls LifeSignal) String() string {
	switch ls {
	case StartLife:
		return "START_LIFE"
	}
	return "unknown signal"
}
