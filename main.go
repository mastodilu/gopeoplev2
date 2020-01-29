package main

import (
	"fmt"

	"github.com/mastodilu/gopeoplev2/person"
)

func main() {
	p := person.NewPerson()
	fmt.Println(p.String())
}
