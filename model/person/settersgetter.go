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

// isEngaged returns true if this Person doesn't have a partner
func (p *Person) isEngaged() bool {
	p.engaged.lock.Lock()
	defer p.engaged.lock.Unlock()
	return p.engaged.value
}

// setEngaged setter
func (p *Person) setEngaged(value bool) {
	p.engaged.lock.Lock()
	defer p.engaged.lock.Unlock()
	p.engaged.value = value
}

// IsFemale returns true if Person is female
func (p *Person) IsFemale() bool {
	return p.sex == 'F'
}
