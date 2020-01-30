package mysignals

// LifeSignal represents a signal
type LifeSignal int

const (
	// StartLife signal sent when a new person is born
	StartLife LifeSignal = iota
	// OneYearOlder signal sent when a person has been alive for another year
	OneYearOlder LifeSignal = iota
	// MaxAgeReached signal sent when a person max age is reached
	MaxAgeReached LifeSignal = iota
	// Stop sends a stop signal, for example to end a person process
	Stop LifeSignal = iota
)

func (ls LifeSignal) String() string {
	switch ls {
	case StartLife:
		return "START_LIFE"
	case OneYearOlder:
		return "ONE_YEAR_OLDER"
	case MaxAgeReached:
		return "MAX_AGE_REACHED"
	}
	return "unknown signal"
}
