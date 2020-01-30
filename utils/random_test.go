package utils

import (
	"testing"
)

func TestNewRandomIntInRange(t *testing.T) {
	// test with (0,0)
	min, max := 0, 0
	got := NewRandomIntInRange(min, max)
	if got < min || got > max {
		t.Errorf("NewRandomIntInRange(%d,%d) = %d, want 0",
			min,
			max,
			got)
	}

	// test with (0,1)
	min, max = 0, 1
	got = NewRandomIntInRange(min, max)
	if got < min || got > max {
		t.Errorf("NewRandomIntInRange(%d,%d) = %d, want a number in the range (%d, %d)",
			min,
			max,
			got,
			min,
			max)
	}

	// test with a wider range of numbers
	for i := 1; i < 10; i++ {
		min, max := 10*i, 12*i
		got := NewRandomIntInRange(min, max)
		if got < min || got > max {
			t.Errorf("NewRandomIntInRange(%d,%d) = %d, want a number in the range (%d, %d)",
				min,
				max,
				got,
				min,
				max)
		}
	}
}
