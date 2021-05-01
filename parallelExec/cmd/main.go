package main

import (
	"fmt"
	"sort"
	"sync"
	"time"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

var WG_MAIN sync.WaitGroup

func showTrackerUnits(units []crtk.Tracker_unit) {
	sort.Slice(units, func(i, j int) bool {
		return units[i].FibVal < units[j].FibVal
	})

	fmt.Println(len(units))
	for _, value := range units {
		if value.Status != nil {
			fmt.Printf("value %d : result is %v\n", value.FibVal, value.Status)
		} else {
			fmt.Printf("value : %d result : %d operations: %d time: %v\n", value.FibVal, value.FinResult, value.Operations, value.ExecutionResTime)
		}
	}
}

func parallelExec(slc []func(v int, ch chan<- crtk.Tracker_unit), parallelCount int, errors int) {

	fmt.Println("Go")

	ch := make(chan crtk.Tracker_unit, parallelCount)
	var resSlc []crtk.Tracker_unit

	// run by bundles
	counterProc := 0
	errCount := 0

	var wg sync.WaitGroup

	// for panic case
	defer func() {
		if r := recover(); r != nil {
			showTrackerUnits(resSlc)
			resSlc = nil
			fmt.Println("Recovered in parallelExec ", r)
			WG_MAIN.Done()
			return
		}
	}()

	for i := 0; i < len(slc); i++ {
		wg.Add(1)
		v := i
		go func(val int) {
			defer wg.Done()
			slc[val](val, ch)
		}(v)
		counterProc++

		if counterProc >= parallelCount || len(slc)-i < parallelCount {
			wg.Wait()
			close(ch)

			for x := range ch {
				if x.Status != nil {
					resSlc = append(resSlc, x)
					errCount++
				} else {
					resSlc = append(resSlc, x)
				}
			}
			if errCount >= errors {
				panic(fmt.Sprintf("Panic from parallelExec err errCount >= errors %d == %d", errCount, errors))
			}
			ch = make(chan crtk.Tracker_unit, parallelCount)
			counterProc = 0
		}

	}
	showTrackerUnits(resSlc)
	WG_MAIN.Done()
}

func main() {
	start := time.Now()
	tasks := crtk.CreateTasks(50)
	WG_MAIN.Add(1)

	parallelExec(tasks, 10, 5)
	fmt.Printf("Execution time : %v\n", time.Since(start))
	WG_MAIN.Wait()
}
