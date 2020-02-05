package loveagency

import (
	"fmt"
	"sync"

	"github.com/mastodilu/gopeoplev2/model/tools/smartphone"
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

var (
	once           sync.Once        // singleton
	receiveRequest chan *PersonInfo // channel where people can send its data
)

// GetInstance returns two channels
// The channel receiveRequest is for sending a registration request. The agency will store
// The data received in this channel and use it to look for potential partner for other people.
func GetInstance() chan<- *PersonInfo {
	once.Do(func() {
		receiveRequest = make(chan *PersonInfo)
		listenForNewCustomers()
	})
	return receiveRequest
}

// listenForNewCustomers listens for new customers on the channel
// and handles the customer list
func listenForNewCustomers() {
	go func() {

		// hashmap of people that asked to be registered as potential partner
		customers := make(map[int]*PersonInfo)

		for {
			newcustomer, ok := <-receiveRequest
			if !ok {
				// channel closed, end this process
				return
			}

			// find a partner for this new person.
			// If a partner is found it will store its key in the hashmap
			id := -1
			for key, customer := range customers {
				if isCompatible(newcustomer, customer) {
					id = key
					break
				}
			}
			if id == -1 {
				// no compatible person was found
				customers[newcustomer.ID()] = newcustomer
			} else {
				// a compatible person was found
				msg := smartphone.NewMessage("love-agency", "a partner was found")
				customers[id].Chat() <- msg
				close(customers[id].Chat())
				delete(customers, id)
			}

		}
	}()
}

// isCompatible returns true if the two curstomers are considered compatible,
// false otherwise
func isCompatible(newcustomer, customer *PersonInfo) bool {
	if newcustomer.Sex() != customer.Sex() {
		fmt.Printf("id:%d is compatible with id:%d\n", newcustomer.ID(), customer.ID())
		return true
	}
	return false
}
