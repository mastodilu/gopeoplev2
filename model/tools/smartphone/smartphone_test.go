package smartphone

import (
	"testing"
	"time"
)

func TestReadNextMessage(t *testing.T) {
	phone := New()

	// chat n 1
	tosend1 := []*Message{
		NewMessage("matteo", "one", nil),
		NewMessage("matteo", "two", nil),
		NewMessage("matteo", "three", nil),
	}

	// chat n 2
	tosend2 := []*Message{
		NewMessage("matteo", "AAA", nil),
		NewMessage("matteo", "BBB", nil),
		NewMessage("matteo", "CCC", nil),
	}

	// chat n 3
	tosend3 := []*Message{
		NewMessage("matteo", "111", nil),
		NewMessage("matteo", "222", nil),
		NewMessage("matteo", "333", nil),
	}

	// testing with 3 writers that all messages are received
	// from first to last in each conversation

	go func() {
		phoneNumber := phone.GiveNumber()
		for _, msg := range tosend1 {
			phoneNumber <- msg
		}
		close(phoneNumber)
	}()
	go func() {
		phoneNumber := phone.GiveNumber()
		for _, msg := range tosend2 {
			phoneNumber <- msg
		}
		close(phoneNumber)
	}()
	go func() {
		phoneNumber := phone.GiveNumber()
		for _, msg := range tosend3 {
			phoneNumber <- msg
		}
		close(phoneNumber)
	}()

	time.Sleep(time.Second * 2)
	totalMsgs := len(tosend1) + len(tosend2) + len(tosend3)

	// testing with 3 writers that all messages are received
	// from first to last in each conversation

	// keep track of the last message received for each conversation
	index1, index2, index3 := 0, 0, 0
	for {
		msg, err := phone.ReadNextMessage()
		if err != nil {
			break // exit loop
		}
		switch msg.Content() {
		case tosend1[index1].Content():
			if index1 < len(tosend1)-1 {
				index1++
			}
		case tosend2[index2].Content():
			if index2 < len(tosend2)-1 {
				index2++
			}
		case tosend3[index3].Content():
			if index3 < len(tosend3)-1 {
				index3++
			}
		}
	}
	if index1 != 2 || index2 != 2 || index3 != 2 {
		t.Errorf("expected to read %d messaged, read %d", totalMsgs, index1+index2+index3)
	}

}
