package geometry

import "math"
import "github.com/programmerjames/go/vector"

type Region struct {
	Faces []Plane
}

func (r Region) Project(v vector.Vector) vector.Vector {
	mDot, mSub := -1.0, vector.Vector{}
	
	for i := 0; i<len(r.Faces); i++ {
		offset := vector.Subtract(v, r.Faces[i].Offset)
		subtract := vector.Vector{
			offset.X * math.Abs(r.Faces[i].Normal.X),
			offset.Y * math.Abs(r.Faces[i].Normal.Y),
			offset.Z * math.Abs(r.Faces[i].Normal.Z),
		}
		
		offset.Multiply(1.0/offset.Magnitude())
		
		dot := vector.Dot(r.Faces[i].Normal, offset)
		
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