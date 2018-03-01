package geometry

import "github.com/programmerjames/go/3d/vector"

type Shape interface {
	Project(vector.Vector) vector.Vector
}
