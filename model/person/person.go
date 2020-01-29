package person

import (
	"fmt"

	"github.com/mastodilu/gopeoplev2/utils"
)

var internalCounter int

// Person represents a person
type Person struct {
	id  int
	age int
	sex byte // 'M' or 'F'
}

// ID person's id getter
func (p *Person) ID() int {
	return p.id
}

// Age person's age getter
func (p *Person) Age() int {
	return p.age
}

// Sex person's sex getter
func (p *Person) Sex() byte {
	return p.sex
}

// New creates and returs a new Person
func New() Person {

	return Person{
		id:  newPersonID(),
		age: 0,
		// sex is M or F
		sex: func() byte {
			if utils.NewRandomIntInRange(0, 1) == 0 {
				return 'M'
			}
			return 'F'
		}(),
	}
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
