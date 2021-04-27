package main

import (
	"fmt"
	"time"

	create_task "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

func parallelExec(slc []func(v int) (int, error)) {
	fmt.Println("Go")
	val := 0
	for _, fu := range slc {
		go fu(val)

		val++
	}

}

func main() {
	tasks := create_task.CreateTask(10)
	parallelExec(tasks)
	time.Sleep(30 * time.Second)
}
