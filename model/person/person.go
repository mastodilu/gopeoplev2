package person

import (
	"fmt"
	"sync"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
	"github.com/mastodilu/gopeoplev2/utils"
)

var internalCounter int

// Age represents how old is a person and contains the semaphore
// to handle mutex on Age.value
type Age struct {
	value int
	lock  sync.Mutex
}

// Person represents a person
type Person struct {
	id       int
	age      Age
	sex      byte // 'M' or 'F'
	lifemsgs chan mysignals.LifeSignal
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

// New creates and returs a new Person
func New(lms chan mysignals.LifeSignal) *Person {
	p := Person{
		id: newPersonID(),
		age: Age{
			value: 0,
			lock:  sync.Mutex{},
		},
		// sex is M or F
		sex: func() byte {
			if utils.NewRandomIntInRange(0, 1) == 0 {
				return 'M'
			}
			return 'F'
		}(),
		lifemsgs: lms,
	}
	return &p
}

// ListenForSignals begin listening for signals (begin living and counting years)
func (p *Person) ListenForSignals() {

	// start to count the years
	go func(ch chan<- mysignals.LifeSignal) {
		for {
			time.Sleep(lifetimings.Year)
			ch <- mysignals.OneYearOlder
			fmt.Println("ping")
		}
	}(p.lifemsgs)

	// handle signals
	go func(ch <-chan mysignals.LifeSignal) {
		for {
			msgin, ok := <-ch
			if !ok {
				fmt.Println("Life closed this channel!")
				break
			}

			switch msgin {
			case mysignals.StartLife:
				fmt.Println("Hello world")
			case mysignals.OneYearOlder:
				p.oneYearOlder()
				fmt.Println("pong")
			}

		}
	}(p.lifemsgs)
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

// oneYearOlder adds one year to the person age
func (p *Person) oneYearOlder() {
	p.age.lock.Lock()
	defer p.age.lock.Unlock()
	p.age.value++
}
