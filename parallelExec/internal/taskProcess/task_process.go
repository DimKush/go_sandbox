package taskProcess

import (
	"fmt"
	"sort"
	"sync"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

func ShowTrackerUnits(units []crtk.Tracker_unit) {
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

func ParallelExec(wg_main *sync.WaitGroup, slc []func(v int, ch chan<- crtk.Tracker_unit), parallelCount int, errors int) {
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
			ShowTrackerUnits(resSlc)
			resSlc = nil
			fmt.Println("Recovered in parallelExec ", r)
			wg_main.Done()
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
	ShowTrackerUnits(resSlc)
	wg_main.Done()
}
