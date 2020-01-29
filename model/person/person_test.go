package person

import "testing"

func TestNew(t *testing.T) {
	newperson := New() // crate a new Person
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
