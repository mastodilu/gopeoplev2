package main

import (
	"fmt"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/loveagency"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
	"github.com/mastodilu/gopeoplev2/model/person"
)

func main() {

	nPpl := 10
	ch := make(chan mysignals.LifeSignal)
	loveAgency := loveagency.GetInstance()

	for i := 0; i < nPpl; i++ {
		p := person.New(ch)
		fmt.Println(p.String())
		ch <- mysignals.StartLife
		loveAgency <- loveagency.NewPersonInfo(p.ID(), p.Sex(), p.Chat())

	}

	time.Sleep(lifetimings.Year * 4)
}
