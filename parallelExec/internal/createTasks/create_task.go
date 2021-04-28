package create_task

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CreateTask(TasksCount int) []func(v int, progress_Tracker *Tracker) (int, error) {
	var taskSlice []func(v int, progress_Tracker *Tracker) (int, error)

	for i := 0; i <= TasksCount; i++ {
		taskSlice = append(taskSlice, func(v int, progress_Tracker *Tracker) (int, error) {
			// get a state of fibVal to show state on screen
			// in this function v variable was send as a copy, don't worry about the data mutation:)

		})
	}

	return taskSlice
}

func showProcess(num int) {
	str := "Process fib of number :" + strconv.Itoa(num)

	var strBuf strings.Builder

	for {
		strBuf.Reset()
		for i := 0; i < 100; i++ {
			strBuf.WriteRune('#')
			fmt.Printf("%s %s\n", str, strBuf.String())
			time.Sleep(1 * time.Second)
		}
	}
}

func countFib(val int) int {
	if val == 0 {
		return 0
	}

	if val < 2 {
		return val
	} else {
		return countFib(val-1) - countFib(val-2)
	}

}
