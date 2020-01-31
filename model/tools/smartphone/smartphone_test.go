package smartphone

import (
	"testing"
	"time"
)

func TestReadNextMessage(t *testing.T) {
	phone := New()
	expected := []*Message{
		NewMessage("matteo", "one"),
		NewMessage("matteo", "two"),
		NewMessage("matteo", "three"),
		NewMessage("matteo", "four"),
		NewMessage("matteo", "five"),
	}
	tosend := []*Message{
		NewMessage("matteo", "one"),
		NewMessage("matteo", "two"),
		NewMessage("matteo", "three"),
		NewMessage("matteo", "four"),
		NewMessage("matteo", "five"),
	}

	go func() {
		phoneNumber := phone.GiveNumber()
		for i := 0; i < len(tosend); i++ {
			phoneNumber <- tosend[i]
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < len(expected); i++ {
		msg, _ := phone.ReadNextMessage()
		if msg.From() != expected[i].From() ||
			msg.Content() != expected[i].Content() {
			t.Errorf("expected %s, got %s", expected[i].String(), msg.String())
		}
	}

}
