package geometry

import "github.com/programmerjames/go/vector"

type Cylinder struct {}

func (c Cylinder) Project(v vector.Vector) vector.Vector {
	return v
}