package taskProcess

import (
	"sync"
	"testing"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

func BenchmarkParallelExec(b *testing.B) {
	var wg_main sync.WaitGroup
	b.ResetTimer()
	tasks := crtk.CreateTasks(10)
	wg_main.Add(1)
	units, _ := ParallelExec(&wg_main, tasks, 10, 5)
	ShowTrackerUnits(units)

	wg_main.Wait()
}
