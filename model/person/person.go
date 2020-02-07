package person

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/loveagency"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
	"github.com/mastodilu/gopeoplev2/model/person/decisionmaker/brain"
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

// engaged is a struct with value=true when Person is engaged, false otherwire
// It contains a boolen value and its sempaphore to make it thread safe
type engaged struct {
	value bool
	lock  sync.Mutex
}

// Person represents a person
type Person struct {
	id         int
	age        Age
	sex        byte                      // 'M' or 'F'
	lifemsgs   chan mysignals.LifeSignal // channel used to handle signals
	smartphone *smartphone.Smartphone    // channel used to handle chats with potential partners and with the agency
	brain      DecisionMaker
	engaged    engaged // evaluated when a partner is found
}

var (
	loveAg = loveagency.GetInstance()
)

// New creates and returs a new Person
func New(lms chan mysignals.LifeSignal) *Person {
	// create new Person
	p := Person{
		id: newPersonID(),
		age: Age{
			value:  15,
			lock:   sync.Mutex{},
			maxage: 100,
		},
		smartphone: smartphone.New(),
		lifemsgs:   lms,
		engaged: engaged{
			value: false,
			lock:  sync.Mutex{},
		},
		// use &brain.Brain{}
		// instead of brain.Brain{}
		// because Brain implements the interface DecisionMaker{} using a pointer receiver
		// 			(b* Brain)Method(...)
		// instead of a value receiver
		// 			(b Brain)Method(...)
		brain: &brain.Brain{},
		// sex is M or F
		sex: func() byte {
			if utils.NewRandomIntInRange(0, 1) == 0 {
				return 'M'
			}
			return 'F'
		}(),
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
			stayInLoop = false
			break
		}

		switch msgin {
		case mysignals.Stop:
			return // end person process
		case mysignals.OneYearOlder:
			p.oneYearOlder()
			// if single, every 2 years, contact the love agency
			if !p.isEngaged() && p.Age()%2 == 0 {
				loveagency.GetInstance() <- loveagency.NewPersonInfo(
					p.ID(),
					p.Sex(),
					p.Chat(),
				)
			}
		case mysignals.MaxAgeReached:
			stayInLoop = false
		}
	}

	fmt.Println("Bye")

}

// handleMessages reads and handles the next message received
func (p *Person) handleMessages() {
	for {
		msg, err := p.smartphone.ReadNextMessage()
		if err != nil {
			// if no messages then check every 6 months
			time.Sleep(lifetimings.Month * 6)
		} else {
			switch msg.From() {
			case "love-agency": // receive a message from love agency because a potential partner was found
				fmt.Printf("%d msg in from love-agency\n", p.ID())
				if p.isEngaged() {
					// do nothing because p has a partner
					// just ignore this message from the love agency
				} else {
					// prevent this Person from contacting the love-agency again
					// during this conversation with a potential partner
					p.setEngaged(true)
					creepymsg := smartphone.NewMessage(
						"stranger",
						"hey, are you single?",
						p.Chat(),
					)
					msg.Contact() <- creepymsg
					log.Printf("stranger %2d: hey %p are you single?\n", p.ID(), msg.Contact())
				}
			case "stranger":
				fmt.Printf("%d msg in from stranger\n", p.ID())
				if p.isEngaged() {
					msg.Contact() <- smartphone.NewMessage(
						"partner",
						"no", // send "no, I'm engaged"
						nil,  // nil because there won't be any further communication
					)
				} else {

					p.setEngaged(true)
					msg.Contact() <- smartphone.NewMessage(
						"partner",
						"yes", // send "yes, I'm single"
						nil,   // nil because there won't be any further communication
					)
					log.Printf("partner %d - %p, yes\n", p.ID(), p.Chat())
				}

			case "partner":
				fmt.Printf("%d msg in from partner\n", p.ID())
				if msg.Content() == "yes" {
					log.Printf("Hurray, id %d is engaged\n", p.ID())
				} else {
					p.setEngaged(false)
					log.Println("Damn, we're not engaged")
				}

			default:
				// ignore for now
				// this message is from an unknown sender
			}
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

// isRightPartner calls IsRightPartner on the interface DecisionMaker
func (p *Person) isRightPartner(p2 *Person) bool {
	// TODO return p.brain.IsRightPartner(intelligent1, intelligent2, pretty1, pretty2 int, sex1, sex2 byte)
	return p.Sex() != p2.Sex()
}
