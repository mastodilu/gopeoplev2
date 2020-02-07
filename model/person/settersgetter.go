package person

import (
	"fmt"

	"github.com/mastodilu/gopeoplev2/model/tools/smartphone"
)

// internalCounter keeps track of used IDs
var internalCounter int

// newPersonID returns the next valid ID
func newPersonID() int {
	internalCounter++
	return internalCounter
}

// ID person's id getter
func (p *Person) ID() int {
	return p.id
}

// Age person's age getter
func (p *Person) Age() int {
	p.age.lock.Lock()
	defer p.age.lock.Unlock()
	value := p.age.value

	return value
}

// Sex person's sex getter
func (p *Person) Sex() byte {
	return p.sex
}

// String returns the Person as formatted string
func (p *Person) String() string {
	return fmt.Sprintf("id:%d age:%d sex:%c", p.ID(), p.Age(), p.Sex())
}

// Chat returns a channel where to write directly to this person
func (p *Person) Chat() chan<- *smartphone.Message {
	return p.smartphone.GiveNumber()
}

// isSingle returns true if this Person doesn't have a partner
func (p *Person) isSingle() bool {
	p.isEngaged.lock.Lock()
	defer p.isEngaged.lock.Unlock()
	return p.isEngaged.value == false
}

// engaged setter
func (p *Person) engaged() {
	p.isEngaged.lock.Lock()
	defer p.isEngaged.lock.Unlock()
	p.isEngaged.value = true
}

// notEngaged setter
func (p *Person) notEngaged() {
	p.isEngaged.lock.Lock()
	defer p.isEngaged.lock.Unlock()
	p.isEngaged.value = false
}
