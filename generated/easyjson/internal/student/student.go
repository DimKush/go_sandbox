package student

type Discipline struct {
	Discipline string
	Mark       int
}

//go:generate easyjson -all student.go
type Student struct {
	FirstName string
	LastName  string
	Age       int
	Marks     []Discipline
}
