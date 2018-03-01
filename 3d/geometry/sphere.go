package geometry

import "github.com/programmerjames/go/3d/vector"

type Sphere struct {}

func (s Sphere) Project(v vector.Vector) vector.Vector {
	return vector.Multiply(v, 1.0/v.Magnitude())
}
