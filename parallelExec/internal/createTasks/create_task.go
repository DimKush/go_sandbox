package create_task

import (
	"sync"
	"time"
)

type Tracker_unit struct {
	FibVal           int
	ExecutionResTime time.Duration
	FinResult        int
	Operations       int
}

func CreateTask(TasksCount int) []func(v int, ch chan<- Tracker_unit, wg *sync.WaitGroup) {
	var taskSlice []func(v int, ch chan<- Tracker_unit, wg *sync.WaitGroup)

	for i := 1; i <= TasksCount; i++ {

		taskSlice = append(taskSlice, func(v int, ch chan<- Tracker_unit, wg *sync.WaitGroup) {
			defer wg.Done()
			start := time.Now()

			oprs := 0
			resVal := countFib(v, &oprs)

			ch <- Tracker_unit{FibVal: v, ExecutionResTime: time.Since(start), FinResult: resVal, Operations: oprs}

		})
	}
	return taskSlice
}

func countFib(val int, operations *int) int {
	*operations += 1

	if val == 0 {
		return 0
	}

	if val < 2 {
		return val
	} else {
		return countFib(val-1, operations) + countFib(val-2, operations)
	}

}
