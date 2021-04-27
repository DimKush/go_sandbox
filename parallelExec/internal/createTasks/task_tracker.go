package create_task

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type tracker_unit struct {
	fibVal   int
	prcTrack string
	prcCount int
}

func (d *tracker_unit) update() {
	if d.prcCount > 20 {
		d.prcTrack = ""
		d.prcCount = 0
	}

	d.prcTrack = d.prcTrack + string('#')
	d.prcCount++
}

type tracker struct {
	log_units []tracker_unit
}

func (d *tracker) show() {
	var mutex = &sync.Mutex{}

	mutex.Lock()
	for _, v := range d.log_units {
		str := "Process fib of number : " + strconv.Itoa(v.fibVal)

		fmt.Printf("\r%s", str)
		time.Sleep(5 * time.Millisecond)

	}
	mutex.Unlock()
}
