package main

import (
	"fmt"
	"sort"
	"sync"
	"time"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

var WG_MAIN sync.WaitGroup

func parallelExec(slc []func(v int, ch chan<- crtk.Task_Unit, wg *sync.WaitGroup), parallelCount int, errors int) {
	var wg sync.WaitGroup
	fmt.Println("Go")
	val := 0

	ch := make(chan crtk.Task_Unit, len(slc))
	fmt.Println(len(slc))
	for _, fu := range slc {
		wg.Add(1)
		go fu(val+1, ch, &wg)

		val++
	}

	wg.Wait()
	//close(ch)

	var resSlc []crtk.Task_Unit
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
	tasks := crtk.CreateTask(49)
	WG_MAIN.Add(1)

	parallelExec(tasks, 10, 3)
	fmt.Printf("Execution time : %v\n", time.Since(start))
	WG_MAIN.Wait()
}
