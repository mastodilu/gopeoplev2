package person

import (
	"testing"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
)

func TestNew(t *testing.T) {
	ch := make(chan mysignals.LifeSignal)

	newperson := New(ch) // crate a new Person
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

func TestListenForSignals(t *testing.T) {
	ch := make(chan mysignals.LifeSignal)
	p := New(ch)
	p.ListenForSignals()
	time.Sleep(lifetimings.Year * 5)
	personAge := p.Age()
	if personAge != 4 && personAge != 5 && personAge != 6 {
		t.Errorf("p.Age() = %d, expected value in range (4,6)", personAge)
	}
}
