package student

type Descilpine struct {
	Descilpine string
	Mark       int
}

//go:generate easyjson -all student.go
type Student struct {
	FirstName string
	LastName  string
	Age       int
	Marks     []Descilpine
}
