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
	ch := make(chan mysignals.LifeSignal)

	p1 := person.New(ch)
	p2 := person.New(ch)
	p3 := person.New(ch)
	p4 := person.New(ch)
	p5 := person.New(ch)

	fmt.Println(p1.String())
	fmt.Println(p2.String())
	fmt.Println(p3.String())
	fmt.Println(p4.String())
	fmt.Println(p5.String())

	time.Sleep(time.Second)

	ch <- mysignals.StartLife
	ch <- mysignals.StartLife
	ch <- mysignals.StartLife
	ch <- mysignals.StartLife
	ch <- mysignals.StartLife

	loveAgency := loveagency.GetInstance()
	loveAgency <- p1
	loveAgency <- p2
	loveAgency <- p3
	loveAgency <- p4
	loveAgency <- p5

	time.Sleep(lifetimings.Year * 4)
}
