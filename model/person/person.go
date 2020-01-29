package person

import (
	"fmt"
	"sync"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
	"github.com/mastodilu/gopeoplev2/utils"
)

// Age represents how old is a person and contains the semaphore
// to handle mutex on Age.value
type Age struct {
	value  int
	lock   sync.Mutex
	maxage int
}

// Person represents a person
type Person struct {
	id       int
	age      Age
	sex      byte // 'M' or 'F'
	lifemsgs chan mysignals.LifeSignal
}

// New creates and returs a new Person
func New(lms chan mysignals.LifeSignal) *Person {
	p := Person{
		id: newPersonID(),
		age: Age{
			value:  0,
			lock:   sync.Mutex{},
			maxage: 100,
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
			// when max age is reached this person must stop getting older
			if p.Age() == p.age.maxage {
				ch <- mysignals.MaxAgeReached
			}

			time.Sleep(lifetimings.Year) // wait one year
			ch <- mysignals.OneYearOlder // signal that one year is passed
		}
	}(p.lifemsgs)

	// handle signals
	go func(ch <-chan mysignals.LifeSignal) {
		continueLife := true
		for continueLife {
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
			case mysignals.MaxAgeReached:
				continueLife = false
			}
		}
		fmt.Println("Bye")
	}(p.lifemsgs)
}

// oneYearOlder adds one year to the person age
func (p *Person) oneYearOlder() {
	p.age.lock.Lock()
	defer p.age.lock.Unlock()
	p.age.value++
}
