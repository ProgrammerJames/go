package geometry

import "math"
import "github.com/programmerjames/go/vector"

type Region struct {
	faces []Plane
}

func NewRegion(faces []Plane) Region {
	return Region{faces}
}

func (r Region) Project(v vector.Vector) vector.Vector {
	mDot, mSub := -1.0, vector.Vector{}
	
	for i := 0; i<len(r.faces); i++ {
		offset := vector.Subtract(v, r.faces[i].Offset)
		subtract := vector.Vector{
			offset.X * math.Abs(r.faces[i].Normal.X),
			offset.Y * math.Abs(r.faces[i].Normal.Y),
			offset.Z * math.Abs(r.faces[i].Normal.Z),
		}
		
		offset.Multiply(1.0/offset.Magnitude())
		
		dot := vector.Dot(r.faces[i].Normal, offset)
		
		if (dot > mDot) {
			mDot = dot
			mSub = subtract
		}
		
		if (dot > 0) {
			v.Subtract(subtract)
		}
	}
	
	if (mDot < 0) {
		v.Subtract(mSub)
	}
	
	return v
}
