package smartphone

import (
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

// New creates a new Message object
func New(from, content string) *Message {
	return &Message{
		from:    from,
		content: content,
		created: time.Now(),
	}
}
