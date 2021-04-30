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

	ch := make(chan crtk.Tracker_unit, len(slc))
	fmt.Println(len(slc))

	var wg sync.WaitGroup
	for i := 0; i < len(slc); i++ {
		wg.Add(1)

		go func(f func(v int, ch chan<- crtk.Tracker_unit)) {
			v := i
			defer wg.Done()
			f(v, ch)
		}(slc[i])

	}
	wg.Wait()

	close(ch)

	var resSlc []crtk.Tracker_unit
	for x := range ch {
		resSlc = append(resSlc, x)
	}
	sort.Slice(resSlc, func(i, j int) bool {
		return resSlc[i].FibVal < resSlc[j].FibVal
	})

	for _, v := range resSlc {
		fmt.Printf("value : %d result : %d operations: %d time: %v\n", v.FibVal, v.FinResult, v.Operations, v.ExecutionResTime)
	}

	WG_MAIN.Done()
}

func main() {
	start := time.Now()
	tasks := crtk.CreateTask(45)
	WG_MAIN.Add(1)

	parallelExec(tasks, 10, 3)
	fmt.Printf("Execution time : %v\n", time.Since(start))
	WG_MAIN.Wait()
}
