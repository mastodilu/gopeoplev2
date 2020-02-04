package person

import (
	"fmt"
	"sync"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
	"github.com/mastodilu/gopeoplev2/model/tools/smartphone"
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
	id         int
	age        Age
	sex        byte                      // 'M' or 'F'
	lifemsgs   chan mysignals.LifeSignal // channel used to handle signals
	smartphone *smartphone.Smartphone    // channel used to handle chats with potential partners and with the agency
}

// New creates and returs a new Person
func New(lms chan mysignals.LifeSignal) *Person {
	// create new Person
	p := Person{
		id: newPersonID(),
		age: Age{
			value:  0,
			lock:   sync.Mutex{},
			maxage: 100,
		},
		smartphone: smartphone.New(),
		// sex is M or F
		sex: func() byte {
			if utils.NewRandomIntInRange(0, 1) == 0 {
				return 'M'
			}
			return 'F'
		}(),
		lifemsgs: lms,
	}
	// start listening for signals
	go (&p).listenForSignals()
	// return Person information
	return &p
}

// ListenForSignals begin listening for signals (begin living and counting years)
func (p *Person) listenForSignals() {
	// handle signals:
	// use a Closure to define the channel as read only
	func(ch <-chan mysignals.LifeSignal) {
		// stay in this loop until StartLife signal
		stayInLoop := true
		for stayInLoop {
			msg := <-ch
			switch msg {
			case mysignals.Stop:
				return // end person process
			case mysignals.StartLife:
				// start to count the years
				go p.begingAging(p.lifemsgs)
				stayInLoop = false
			}
		}

		// handle signals
		stayInLoop = true
		for stayInLoop {

			msgin, ok := <-ch
			if !ok {
				fmt.Println("Life closed this channel!")
				break
			}

			switch msgin {
			case mysignals.Stop:
				return // end person process
			case mysignals.OneYearOlder:
				p.oneYearOlder()
				fmt.Println("age is", p.Age())
			case mysignals.MaxAgeReached:
				stayInLoop = false
			default:
				// TODO p.readMessages()
			}
		}

		fmt.Println("Bye")
	}(p.lifemsgs)
}

// ReadMessages reads and handles the next message received
func (p *Person) ReadMessages() {
	// read message
	// handle message
}

// oneYearOlder adds one year to the person age
func (p *Person) oneYearOlder() {
	p.age.lock.Lock()
	defer p.age.lock.Unlock()
	p.age.value++
}

// begingAging starts to increase the age of the given person
func (p *Person) begingAging(ch chan<- mysignals.LifeSignal) {
	for {
		// when max age is reached this person must stop getting older
		if p.Age() >= p.age.maxage {
			ch <- mysignals.MaxAgeReached
			return // exit loop: stop counting the years
		}

		time.Sleep(lifetimings.Year) // wait one year
		ch <- mysignals.OneYearOlder // signal that one year is passed
	}
}
