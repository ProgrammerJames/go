package geometry

import "github.com/programmerjames/go/vector"

type Sphere struct {}

func (s Shpere) Project(v vector.Vector) vector.Vector {
	return vector.Multiply(v, 1.0/v.Magnitude())
}