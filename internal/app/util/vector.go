package util

type Vector struct {
	X, Y int32
}

var (
	VectorUp    = Vector{0, -1} // Since Y grows down
	VectorDown  = Vector{0, 1}
	VectorLeft  = Vector{-1, 0}
	VectorRight = Vector{1, 0}
	VectorZero  = Vector{0, 0}
)

func (v *Vector) Sum(other Vector) Vector {
	if other == VectorZero {
		return *v
	}
	return Vector{v.X + other.X, v.Y + other.Y}
}

func (v *Vector) MulScalar(s int32) Vector {
	return Vector{v.X * s, v.Y * s}
}

func NewVector(x, y int32) Vector {
	return Vector{x, y}
}
