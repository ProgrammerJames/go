package geometry

import "math"
import "github.com/programmerjames/go/vector"

type Plane struct {
	Offset, Normal vector.Vector
}

func (p Plane) Project(v vector.Vector) vector.Vector {
	offset := vector.Subtract(v, p.Offset)
	sub := vector.Vector{offset.X * math.Abs(p.Normal.X), offset.Y * math.Abs(p.Normal.Y), offset.Z * math.Abs(p.Normal.Z)}
	
	return vector.Subtract(v, sub)
}
