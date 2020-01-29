package mysignals

// LifeSignal represents a signal
type LifeSignal int

const (
	_ = iota // 0
	// StartLife signal sent when a new person is born
	StartLife LifeSignal = iota // 1
	// OneYearOlder signal sent when a person has been alive for another year
	OneYearOlder LifeSignal = iota // 2
)

func (ls LifeSignal) String() string {
	switch ls {
	case StartLife:
		return "START_LIFE"
	case OneYearOlder:
		return "ONE_YEAR_OLDER"
	}
	return "unknown signal"
}
