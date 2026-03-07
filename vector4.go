package kmath

import (
	"fmt"
	"math"
)

type Vec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func MakeVec4(x float64, y float64, z float64, w float64) Vec4 {
	return Vec4{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func (v *Vec4) GetLength() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(v.W, 2))
}

func (v Vec4) Normalize() Vec4 {
	return Vec4{
		X: v.X / v.GetLength(),
		Y: v.Y / v.GetLength(),
		Z: v.Z / v.GetLength(),
		W: v.W / v.GetLength(),
	}
}

func (v Vec4) Add(other Vec4) Vec4 {
	return Vec4{v.X + other.X, v.Y + other.Y, v.Z + other.Z, v.W + other.W}
}

func (v Vec4) Sub(other Vec4) Vec4 {
	return Vec4{v.X - other.X, v.Y - other.Y, v.Z - other.Z, v.W - other.W}
}

func (v Vec4) Scale(other Vec4) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

func (v Vec4) DivScale(scale float64) Vec4 {
	return Vec4{X: v.X / scale, Y: v.Y / scale, Z: v.Z / scale, W: v.W / scale}
}

func (v Vec4) Print() {
	fmt.Printf("x : %v, y : %v z : %v, w : %v \n", v.X, v.Y, v.Z, v.W)
}
