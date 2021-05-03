package create_task_test

import (
	"testing"

	crtk "github.com/DimKush/go_sandbox/tree/main/parallelExec/internal/createTasks"
)

func TestCreateTasks(t *testing.T) {
	countTasks := 10
	var checkSlice []func(v int, ch chan<- crtk.Tracker_unit)

	for i := 0; i <= countTasks; i++ {
		checkSlice = append(checkSlice, func(v int, ch chan<- crtk.Tracker_unit) {})
	}

	testSlice := crtk.CreateTasks(countTasks)

	if len(checkSlice) != len(testSlice) {
		t.Fatalf("Fatal test in TestCreateTasks len: expected %d fact %d", len(checkSlice), len(testSlice))
	}
}
