package create_task

import (
	"errors"
	"fmt"
	"strings"
)

func CreateTask(sizeTasks int) []func() {
	var taskSlice []func()

	for i := 0; i <= sizeTasks; i++ {
		taskSlice = append(taskSlice, func(v int) int, error{})
	}

}

func showProcess(num int) {
	fmt.Printf("\n Process fib of number(%d):", num)

	var strBuf strings.Builder

	for {
		strBuf.Reset()
		for i := 0; i < 10; i++ {
			strBuf.WriteRune('#')
			fmt.Printf("\r%s", strBuf)
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
