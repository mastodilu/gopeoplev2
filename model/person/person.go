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

	// stay in this loop until StartLife signal
	stayInLoop := true
	for stayInLoop {
		msg := <-p.lifemsgs
		if msg == mysignals.StartLife {
			// overwrite this channel so that it becomes "private"
			p.lifemsgs = make(chan mysignals.LifeSignal)
			// start to count the years
			go p.begingAging(p.lifemsgs)
			stayInLoop = false
		}
	}

	// read messages
	go p.handleMessages()

	// handle signals
	stayInLoop = true
	for stayInLoop {

		msgin, ok := <-p.lifemsgs
		if !ok {
			fmt.Println("Life closed this channel!")
			break
		}

		switch msgin {
		case mysignals.Stop:
			return // end person process
		case mysignals.OneYearOlder:
			p.oneYearOlder()
		case mysignals.MaxAgeReached:
			stayInLoop = false
		}
	}

	fmt.Println("Bye")

}

// handleMessages reads and handles the next message received
func (p *Person) handleMessages() {
	fmt.Println(p.id, "is waiting for new messages")
	for {
		msg, err := p.smartphone.ReadNextMessage()
		if err != nil {
			// if no messages then check every 6 months
			time.Sleep(lifetimings.Month * 6)
		} else {
			fmt.Println(msg.From())
		}
	}
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
