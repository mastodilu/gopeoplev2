package smartphone

import (
	"fmt"
	"sync"
)

// Smartphone -
type Smartphone struct {
	ch   chan *Message
	msgs []*Message
	lock sync.Mutex
}

// New creates new smartphone
func New() *Smartphone {
	phone := Smartphone{
		ch: make(chan *Message),
	}
	phone.receiveMessages()
	return &phone
}

// GiveNumber creates a middleware that forwards to the smartphone channel
// messages received from different sources
func (s *Smartphone) GiveNumber() chan<- Message {
	newch := make(chan Message)
	go func() {
		for {
			msg, ok := <-newch
			if !ok {
				return // channel closed
			}
			s.ch <- &msg // forward message to actual smartphone
		}
	}()
	return newch
}

// receiveMessages always listen for incoming messages to store
func (s *Smartphone) receiveMessages() {
	go func() {
		for {
			msg, ok := <-s.ch
			if !ok {
				return // main channel closed
			}
			// using a function to wrap its content only
			// to defer sync.Mutex Unlock method
			func() {
				s.lock.Lock()
				defer s.lock.Unlock()
				s.msgs = append(s.msgs, msg)
			}()
		}
	}()
}

// ReadNextMessage returns the oldest message in the list
// or an error if no message is available
func (s *Smartphone) ReadNextMessage() (*Message, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.msgs) < 1 {
		return &Message{}, fmt.Errorf("you have no new message")
	}
	msg := s.msgs[0]
	s.msgs = s.msgs[1:]
	return msg, nil
}
