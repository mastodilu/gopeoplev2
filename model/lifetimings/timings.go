package lifetimings

import (
	"time"
)

const (
	// Year represents one Year
	Year = time.Second * 3 // one year passes every N seconds
	// Month represents one month
	Month = Year / 12
)
