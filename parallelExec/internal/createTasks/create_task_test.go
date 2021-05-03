package create_task

import (
	"testing"
)

func TestCreateTasks(t *testing.T) {
	countTasks := 10
	var checkSlice []func(v int, ch chan<- Tracker_unit)

	for i := 0; i <= countTasks; i++ {
		checkSlice = append(checkSlice, func(v int, ch chan<- Tracker_unit) {})
	}

	testSlice := CreateTasks(countTasks)

	if len(checkSlice) != len(testSlice) {
		t.Fatalf("Fatal test in TestCreateTasks len: expected %d fact %d", len(checkSlice), len(testSlice))
	}
}

func TestCreateTask(t *testing.T) {
	testedTask := CreateTask()
	if testedTask == nil {
		t.Fatalf("Fatal test in TestCreateTask")
	}

}

func TestCountFib(t *testing.T) {
	var operations int
	val := 10

	resExpected := 55
	operationsExpected := 177
	res := countFib(val, &operations)

	if res != resExpected {
		t.Fatalf("Fatal test in TestCountFib res expected : %d fact : %d", resExpected, res)
	}

	if operations != operationsExpected {
		t.Fatalf("Fatal test in TestCountFib operations expected : %d fact : %d", operationsExpected, operations)
	}

}
