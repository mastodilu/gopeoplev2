package person

import "fmt"

var internalCounter int

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

// newPersonID returns the next valid ID
func newPersonID() int {
	internalCounter++
	return internalCounter
}

// Smartphone returns a channel where to write directly to this person
func (p *Person) Smartphone() chan<- *Person {
	return p.smartphone
}
