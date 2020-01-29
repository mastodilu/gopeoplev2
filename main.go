package main

import (
	"fmt"

	"github.com/mastodilu/gopeoplev2/model"
)

func main() {
	p := model.NewPerson()
	fmt.Println(p.String())
}
