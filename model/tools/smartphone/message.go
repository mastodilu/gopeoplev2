package smartphone

import (
	"fmt"
	"time"
)

// Message represents a message in the smartphone
type Message struct {
	from    string
	content string
	created time.Time
}

// From getter
func (msg *Message) From() string {
	return msg.from
}

// Content getter
func (msg *Message) Content() string {
	return msg.content
}

// Created getter
func (msg *Message) Created() time.Time {
	return msg.created
}

// NewMessage creates a new Message object
func NewMessage(from, content string) *Message {
	return &Message{
		from:    from,
		content: content,
		created: time.Now(),
	}
}

func (msg *Message) String() string {
	return fmt.Sprintf(
		"Message: [From:%s][Content:%s]",
		msg.from,
		msg.content,
	)
}
