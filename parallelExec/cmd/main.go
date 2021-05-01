package main

import (
	"fmt"
	"sort"
	"sync"
	"time"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

var WG_MAIN sync.WaitGroup

func parallelExec(slc []func(v int, ch chan<- crtk.Tracker_unit), parallelCount int, errors int) {

	fmt.Println("Go")

	ch := make(chan crtk.Tracker_unit, parallelCount)
	fmt.Println(len(slc))
	var resSlc []crtk.Tracker_unit

	// run by bundles
	counterProc := 0
	errCount := 0
	counter := 0

	var wg sync.WaitGroup
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
					fmt.Println("err")
					errCount++
				} else {
					counter++
					//fmt.Printf("insert : %d result : %d operations: %d time: %v\n", x.FibVal, x.FinResult, x.Operations, x.ExecutionResTime)
					resSlc = append(resSlc, x)
				}
			}
			if errCount >= errors {
				panic(fmt.Sprintf("Panic from parallelExec err errCount >= errors %d == %d", errCount, errors))
			}

			//fmt.Println(counter, i, counterProc)
			ch = make(chan crtk.Tracker_unit, parallelCount)
			counterProc = 0
		}

	}

	sort.Slice(resSlc, func(i, j int) bool {
		return resSlc[i].FibVal < resSlc[j].FibVal
	})

	for _, v := range resSlc {
		fmt.Printf("value : %d result : %d operations: %d time: %v\n", v.FibVal, v.FinResult, v.Operations, v.ExecutionResTime)
	}

	// for panic case
	defer func() {
		if r := recover(); r != nil {
			close(ch)
			resSlc = nil
			fmt.Println("Recovered in parallelExec", r)
			WG_MAIN.Done()
			return
		}
	}()

	WG_MAIN.Done()
}

func main() {
	start := time.Now()
	tasks := crtk.CreateTask(55)
	WG_MAIN.Add(1)

	parallelExec(tasks, 10, 3)
	fmt.Printf("Execution time : %v\n", time.Since(start))
	WG_MAIN.Wait()
}
