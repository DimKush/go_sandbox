package create_task

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type Tracker_unit struct {
	FibVal           int
	ExecutionResTime time.Duration
	FinResult        int
	Operations       int
}

type Task_Unit struct {
	unit Tracker_unit
	err  error
}

func CreateTask(TasksCount int) []func(v int, ch chan<- Task_Unit, wg *sync.WaitGroup) {
	var taskSlice []func(v int, ch chan<- Task_Unit, wg *sync.WaitGroup)

	for i := 1; i <= TasksCount; i++ {

		taskSlice = append(taskSlice, func(v int, ch chan<- Task_Unit, wg *sync.WaitGroup) {
			defer wg.Done()
			start := time.Now()

			oprs := 0
			resVal := countFib(v, &oprs)

			// if the countFib func count more that 10 seconds
			if time.Since(start) > time.Duration(8*time.Second) {
				str := "To slow for count value : " + strconv.Itoa(v)
				ch <- Task_Unit{unit: Tracker_unit{}, err: errors.New(str)}
			}

			ch <- Task_Unit{unit: Tracker_unit{FibVal: v,
				ExecutionResTime: time.Since(start),
				FinResult:        resVal,
				Operations:       oprs},
				err: nil}
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
