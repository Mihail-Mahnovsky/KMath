package kmath

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X float64
	Y float64
}

func MakeVec2(x float64, y float64) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

func (v *Vec2) GetLength() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v Vec2) Normalize() Vec2 {
	return Vec2{
		X: v.X / v.GetLength(),
		Y: v.Y / v.GetLength(),
	}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Scale(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) DivScale(scale float64) Vec2 {
	return Vec2{X: v.X / scale, Y: v.Y / scale}
}

func (v Vec2) Print() {
	fmt.Printf("x : %v, y : %v \n", v.X, v.Y)
}
