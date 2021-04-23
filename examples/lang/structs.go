package lang

import "github.com/mimic-go/mimic/examples/lang/pk1"

type Simple struct {
	Field1 int
	field1 int
	Field2 float64
	field2 float64
}

type Complex1 struct {
	Simple

	Field3 *pkg1.Complex
}
