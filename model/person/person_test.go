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
	p1 := New(ch)
	p2 := New(ch)
	p3 := New(ch)

	ch <- mysignals.StartLife
	ch <- mysignals.StartLife
	ch <- mysignals.StartLife

	// wait about 3 years
	time.Sleep(lifetimings.Year * 3)

	if p1.Age() != 2 && p1.Age() != 3 && p1.Age() != 4 {
		t.Errorf("p1.Age() = %d, expected value in range (2, 4)", p1.Age())
	}
	if p2.Age() != 2 && p2.Age() != 3 && p2.Age() != 4 {
		t.Errorf("p2.Age() = %d, expected value in range (2, 4)", p1.Age())
	}
	if p3.Age() != 2 && p3.Age() != 3 && p3.Age() != 4 {
		t.Errorf("p3.Age() = %d, expected value in range (2, 4)", p1.Age())
	}

}
