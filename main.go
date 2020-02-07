package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mastodilu/gopeoplev2/model/lifetimings"
	"github.com/mastodilu/gopeoplev2/model/mysignals"
	"github.com/mastodilu/gopeoplev2/model/person"
)

func main() {

	// set the log output file
	logfile, err := os.Create("out.log")
	if err != nil {
		log.Fatalln(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	nPpl := 5
	ch := make(chan mysignals.LifeSignal)
	defer close(ch)

	for i := 0; i < nPpl; i++ {
		p := person.New(ch)
		fmt.Println(p.String())
		ch <- mysignals.StartLife

		// loveAgency <- loveagency.NewPersonInfo(
		// 	p.ID(),
		// 	p.Sex(),
		// 	p.Chat(),
		// )

	}

	time.Sleep(lifetimings.Year * 10)
}
