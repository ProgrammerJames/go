package geometry

import "github.com/programmerjames/go/vector"

type Plane struct {
	Offset, Normal vector.Vector
}

func (p *Plane) Project(v vector.Vector) vector.Vector {
	offset := vector.Subtract(v, p.Offset)
	sub := vector.Vector{offset.X * p.Normal.X, offset.Y * p.Normal.Y, offset.Z * p.Normal.Z}
	
	return vector.Subtract(v, sub)
}