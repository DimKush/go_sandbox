package main

import (
	"fmt"

	"github.com/DimKush/go_sandbox/tree/main/generated/easyjson/internal/student"
)

func main() {
	tstJson := `{"FirstName":"John","LastName":"Pertrov","Age":22,"Marks":[{"Discipline":"Math","Mark":3}]}`
	jsonb := []byte(tstJson)
	fmt.Println(jsonb)
	std := &student.Student{}
	std.UnmarshalJSON(jsonb)
	fmt.Println(std)
}
