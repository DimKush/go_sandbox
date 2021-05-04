package taskProcess

import (
	"sync"
	"testing"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

func TestParallelExec(t *testing.T) {
	var unitsCompare []crtk.Tracker_unit

	unitsCompare = append(unitsCompare,
		crtk.Tracker_unit{FibVal: 0, FinResult: 0, Operations: 1})
	unitsCompare = append(unitsCompare,
		crtk.Tracker_unit{FibVal: 1, FinResult: 1, Operations: 1})
	unitsCompare = append(unitsCompare,
		crtk.Tracker_unit{FibVal: 2, FinResult: 1, Operations: 3})
	unitsCompare = append(unitsCompare,
		crtk.Tracker_unit{FibVal: 3, FinResult: 2, Operations: 5})
	unitsCompare = append(unitsCompare,
		crtk.Tracker_unit{FibVal: 4, FinResult: 3, Operations: 9})
	unitsCompare = append(unitsCompare,
		crtk.Tracker_unit{FibVal: 5, FinResult: 5, Operations: 15})

	var wg_main sync.WaitGroup
	tasks := crtk.CreateTasks(5)
	wg_main.Add(1)

	units, _ := ParallelExec(&wg_main, tasks, 10, 5)

	wg_main.Wait()

	for index, _ := range unitsCompare {
		if unitsCompare[index].FibVal != units[index].FibVal {
			t.Errorf("Incorrect FibVal in units[%d]. Expected: %d fact: %d", index, unitsCompare[index].FibVal, units[index].FibVal)
		}
		if unitsCompare[index].FinResult != units[index].FinResult {
			t.Errorf("Incorrect FinResult in units[%d]. Expected: %d fact: %d", index, unitsCompare[index].FinResult, units[index].FinResult)
		}
		if unitsCompare[index].FibVal != units[index].FibVal {
			t.Errorf("Incorrect Operations in units[%d]. Expected: %d fact: %d", index, unitsCompare[index].Operations, units[index].Operations)
		}

	}

}
