package main

import (
	"fmt"

	"github.com/mastodilu/gopeoplev2/model/person"
)

func main() {
	p := person.NewPerson()
	fmt.Println(p.String())
}
