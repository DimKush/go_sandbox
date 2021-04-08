package custom

import (
	"fmt"
	"go_sandbox/calendar"
	"log"
	"time"
)

func print(date *calendar.Date) {
	fmt.Println(date.GetYear(), date.GetMonth(), date.GetDay())
}

func Process() {
	var date calendar.Date

	err := date.SetDay(time.Now().Day())
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(int(time.Now().Month()))
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetYear(time.Now().Year())
	if err != nil {
		log.Fatal(err)
	}

	print(&date)

}
