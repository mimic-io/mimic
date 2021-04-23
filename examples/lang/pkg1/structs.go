package pkg1

type Simple struct {
	Field1 int
	field1 int
	Field2 float64
	field2 float64
}

type Complex1 struct {
	Simple

	Field3 map[string]string
	Field4 []string
	Field5 Simple
	Field6 *Simple
	Field7 []Simple
	Field8 []*Simple
}

type Complex2 struct {
	Complex1

	Field100 Complex1
	Field1001 *Complex1
}




