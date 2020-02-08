package person

import (
	"testing"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
)

func TestNew(t *testing.T) {
	ch := make(chan mysignals.LifeSignal)

	// crate a new Person
	newperson := New(ch)

	if newperson.ID() <= 0 {
		t.Errorf("got %d, expected value > 0", newperson.ID())
	}

	if newperson.Age() != 15 {
		t.Errorf("got %d, expected 15", newperson.Age())
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
	close(ch)

	// wait about 3 years
	time.Sleep(lifetimings.Year * 3)

	if p1.Age() != 17 && p1.Age() != 18 && p1.Age() != 19 {
		t.Errorf("p1.Age() = %d, expected value in range (17, 19)", p1.Age())
	}
	if p2.Age() != 17 && p2.Age() != 18 && p2.Age() != 19 {
		t.Errorf("p2.Age() = %d, expected value in range (17, 19)", p1.Age())
	}
	if p3.Age() != 17 && p3.Age() != 18 && p3.Age() != 19 {
		t.Errorf("p3.Age() = %d, expected value in range (17, 19)", p1.Age())
	}
}
