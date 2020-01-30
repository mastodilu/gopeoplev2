package loveagency

import (
	"sync"

	"github.com/mastodilu/gopeoplev2/model/person"
)

/*
hashmap {
	ID persona --> ultima versione dei dati di una persona
}

- un processo sempre attivo ascolta i messaggi in arrivo e aggiorna la mappa
- il partner viene cercato in questa hashmap

- GetInstance restituisce due channel:
	- uno per le richieste di iscrizione
	- e uno per inviare i potenziali partner trovati
*/

var once sync.Once                     // singleton
var receiveRequest chan *person.Person // channel where people can send its data

// GetInstance returns two channels
// The channel receiveRequest is for sending a registration request. The agency will store
// The data received in this channel and use it to look for potential partner for other people.
func GetInstance() chan<- *person.Person {
	once.Do(func() {
		receiveRequest = make(chan *person.Person)
		// TODO call listenForNewCustomers()
	})
	return receiveRequest
}

// CloseInstance closes the current instance and resets all the variables
func CloseInstance() {
	// end listen() process
	close(receiveRequest)
	once = sync.Once{}
}

// listenForNewCustomers listens for new customers on the channel
// and handles the customer list
func listenForNewCustomers() {
	go func() {
		// hashmap of people that asked to be registered as potential partner
		registeredClients := make(map[int]*person.Person)
		hashmapLock := sync.Mutex{}

		for {
			customer, ok := <-receiveRequest
			if !ok {
				// channel closed, end this process
				return
			}

			// find a partner for this new person.
			// If a partner is found id will store its key in the hashmap
			id := -1
			for key, person := range registeredClients {
				if isCompatible(customer, person) {
					id = key
					break
				}
			}
			if id == -1 {
				// no compatible person was found
				hashmapLock.Lock()
				registeredClients[customer.ID()] = customer // store this person
				hashmapLock.Unlock()
			} else {
				hashmapLock.Lock()
				matched := registeredClients[id]
				delete(registeredClients, id)
				hashmapLock.Unlock()
				customer.Smartphone() <- matched
			}

		}
	}()
}

// isCompatible returns true if the two person are compatible
// returns false if they're not
func isCompatible(customer, person *person.Person) bool {
	return false
}
