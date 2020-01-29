package person

import (
	"testing"

	"github.com/mastodilu/gopeoplev2/model/mysignals"
)

func TestNew(t *testing.T) {
	lifemsgRead := make(<-chan mysignals.LifeSignal)
	lifemsgWrite := make(chan<- mysignals.LifeSignal)

	newperson := New(lifemsgRead, lifemsgWrite) // crate a new Person
	if newperson.ID() <= 0 {
		t.Errorf("got %d, expected value > 0", newperson.ID())
	}

	if newperson.Age() != 0 {
		t.Errorf("got %d, expected 0", newperson.Age())
	}

	if newperson.Sex() != 'M' && newperson.Sex() != 'F' {
		t.Errorf("got %d, expected 'M' or 'F'", newperson.Sex())
	}
}
