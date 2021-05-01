package create_task

import (
	"errors"
	"time"
)

type Tracker_unit struct {
	FibVal           int
	ExecutionResTime time.Duration
	FinResult        int
	Operations       int
	Status           error
}

func CreateTasks(TasksCount int) []func(v int, ch chan<- Tracker_unit) {
	var taskSlice []func(v int, ch chan<- Tracker_unit)

	for i := 0; i <= TasksCount; i++ {

		taskSlice = append(taskSlice, func(v int, ch chan<- Tracker_unit) {
			start := time.Now()

			oprs := 0
			resVal := countFib(v, &oprs)

			// if the countFib func count more that 8 seconds
			if time.Since(start) > time.Duration(8*time.Second) {
				str := "To slow for count value : " + time.Since(start).String()
				ch <- Tracker_unit{FibVal: v, Status: errors.New(str)}
				return
			}

			ch <- Tracker_unit{FibVal: v,
				ExecutionResTime: time.Since(start),
				FinResult:        resVal,
				Operations:       oprs}
		})
	}
	return taskSlice
}

func CreateTask() func(v int, ch chan<- Tracker_unit) {
	task := func(v int, ch chan<- Tracker_unit) {
		start := time.Now()

		oprs := 0
		resVal := countFib(v, &oprs)

		// if the countFib func count more that 8 seconds
		if time.Since(start) > time.Duration(8*time.Second) {
			str := "To slow for count value. time is : " + time.Since(start).String()
			ch <- Tracker_unit{FibVal: v, Status: errors.New(str)}
			return
		}

		ch <- Tracker_unit{FibVal: v,
			ExecutionResTime: time.Since(start),
			FinResult:        resVal,
			Operations:       oprs}
	}

	return task
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
