package kmath

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func MakeVec3(x float64, y float64, z float64) Vec4 {
	return Vec4{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v Vec3) GetLength() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v Vec3) Normalize() Vec3 {
	return Vec3{
		X: v.X / v.GetLength(),
		Y: v.Y / v.GetLength(),
		Z: v.Z / v.GetLength(),
	}
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vec3) Scale(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) DivScale(scale float64) Vec3 {
	return Vec3{X: v.X / scale, Y: v.Y / scale, Z: v.Z / scale}
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Mul(scale float64) Vec3 {
	return Vec3{X: v.X * scale, Y: v.Y * scale, Z: v.Z * scale}
}

func (v Vec3) Distance(other Vec3) float64 {
	return v.Sub(other).GetLength()
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vec3) Lerp(other Vec3, t float64) Vec3 {
	return Vec3{
		X: v.X + (other.X-v.X)*t,
		Y: v.Y + (other.Y-v.Y)*t,
		Z: v.Z + (other.Z-v.Z)*t,
	}
}

func (v Vec3) Equals(other Vec3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

func (v Vec3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

func (v Vec3) Print() {
	fmt.Printf("x : %v, y : %v z : %v \n", v.X, v.Y, v.Z)
}
