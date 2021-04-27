package create_task

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var tracV tracker

func CreateTask(TasksCount int) []func(v int) (int, error) {
	var taskSlice []func(v int) (int, error)

	for i := 0; i <= TasksCount; i++ {
		taskSlice = append(taskSlice, func(v int) (int, error) {
			//tracV.log_units = append(tracV.log_units, tracker_unit{fibVal: v, procTrack: ""})
			//go tracV.show()
			val, err := countFib(v)
			if err != nil {
				return 0, nil
			}

			return val, err
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

func countFib(val int) (int, error) {
	if val == 0 {
		return 0, errors.New("Fib function argument can't be 0")
	}
	if val < 2 {
		return val, nil
	} else {
		v1, _ := countFib(val - 1)
		v2, _ := countFib(val - 2)
		return v1 - v2, nil
	}

}
