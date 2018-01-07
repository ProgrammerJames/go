package vector

import "math"

type Vector struct {
	X, Y, Z float64
}

func (a *Vector) Add(b Vector) {
	a.X, a.Y, a.Z = a.X+b.X, a.Y+b.Y, a.Z+b.Z
}

func (a *Vector) Subtract(b Vector) {
	a.X, a.Y, a.Z = a.X-b.X, a.Y-b.Y, a.Z-b.Z
}

func (v *Vector) Multiply(s float64) {
	v.X, v.Y, v.Z = v.X*s, v.Y*s, v.Z*s
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt((v.X*v.X)+(v.Y*v.Y)+(v.Z*v.Z))
}

func Add(a, b Vector) Vector {
	return Vector{a.X+b.X, a.Y+b.Y, a.Z+b.Z}
}

func Subtract(a, b Vector) Vector {
	return Vector{a.X-b.X, a.Y-b.Y, a.Z-b.Z}
}

func Multiply(a Vector, b float64) Vector {
	return Vector{a.X*b, a.Y*b, a.Z*b}
}

func Rotate(a, b Vector, angle float64) Vector {
	s, c := math.Sin(angle), math.Cos(angle)
	
	return Vector{
		(a.X*(b.X*b.X*(1-c)+c))+(a.Y*(b.X*b.Y*(1-c)-(b.Z*s)))+(a.Z*(b.X*b.Z*(1-c)+(b.Y*s))),
		(a.X*(b.X*b.Y*(1-c)+(b.Z*s)))+(a.Y*(b.Y*b.Y*(1-c)+c))+(a.Z*(b.Y*b.Z*(1-c)-(b.X*s))),
		(a.X*(b.X*b.Z*(1-c)-(b.Y*s)))+(a.Y*(b.Y*b.Z*(1-c)+(b.X*s)))+(a.Z*(b.Z*b.Z*(1-c)+c)),
	}
}

func Cross(a, b Vector) Vector {
	return Vector{(a.Y*b.Z)-(a.Z*b.Y), (a.Z*b.X)-(a.X*b.Z), (a.X*b.Y)-(a.Y*b.X)}
}

func Dot(a, b Vector) float64 {
	return (a.X*b.X)+(a.Y*b.Y)+(a.Z*b.Z)
}