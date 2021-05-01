package main

import (
	"fmt"
	"sync"
	"time"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
	tp "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/taskProcess"
)

func main() {
	var wg_main sync.WaitGroup
	start := time.Now()
	tasks := crtk.CreateTasks(50)
	wg_main.Add(1)

	tp.ParallelExec(&wg_main, tasks, 10, 5)
	fmt.Printf("Execution time : %v\n", time.Since(start))
	wg_main.Wait()
}
