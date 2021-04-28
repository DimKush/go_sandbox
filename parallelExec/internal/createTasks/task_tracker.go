package create_task

import (
	"sort"
	"strings"

	"github.com/buger/goterm"
)

type tracker_unit struct {
	fibVal      int
	fibValAddrs *int
	prcCount    int
}

type Tracker struct {
	log_units []tracker_unit
}

func (d *Tracker) updateValStatus(p *int) {
	idx := sort.Search(len(d.log_units), func(i int) bool {
		return d.log_units[i].fibValAddrs == p
	})

	d.log_units[idx].prcCount += 1
}

func (d *Tracker) Show() {
	goterm.Clear()

	var strBuf strings.Builder
	for _, unit := range d.log_units {
		strBuf.Reset()

		for i := 0; i < unit.prcCount; i++ {
			strBuf.WriteRune('#')
		}

	}
}
