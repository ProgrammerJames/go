package geometry

import "github.com/programmerjames/go/vector"

type Shape interface {
	Project(vector.Vector) vector.Vector
}