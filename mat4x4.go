package kmath

type Mat4x4 struct {
	Buf [4][4]float64
}

func MakeMat4x4(b [4][4]float64) Mat4x4 {
	return Mat4x4{
		Buf: b,
	}
}

func (m4 *Mat4x4) Add(other *Mat4x4) Mat4x4 {
	var newM Mat4x4

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			newM.Buf[y][x] = m4.Buf[y][x] + other.Buf[y][x]
		}
	}

	return newM
}

func (m4 *Mat4x4) Sub(other *Mat4x4) Mat4x4 {
	var newM Mat4x4

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			newM.Buf[y][x] = m4.Buf[y][x] - other.Buf[y][x]
		}
	}

	return newM
}

func (m4 *Mat4x4) Mul(other *Mat4x4) Mat4x4 {
	var newM Mat4x4

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			var sum float64
			for k := 0; k < 4; k++ {
				sum += m4.Buf[y][k] * other.Buf[k][x]
			}
			newM.Buf[y][x] = sum
		}
	}

	return newM
}

func (m *Mat4x4) MulVec3(v Vec3) Vec3 {
	return Vec3{
		X: v.X*m.Buf[0][0] + v.Y*m.Buf[0][1] + v.Z*m.Buf[0][2] + m.Buf[0][3],
		Y: v.X*m.Buf[1][0] + v.Y*m.Buf[1][1] + v.Z*m.Buf[1][2] + m.Buf[1][3],
		Z: v.X*m.Buf[2][0] + v.Y*m.Buf[2][1] + v.Z*m.Buf[2][2] + m.Buf[2][3],
	}
}

func (m *Mat4x4) MulVec4(v Vec4) Vec4 {
	return Vec4{
		X: v.X*m.Buf[0][0] + v.Y*m.Buf[0][1] + v.Z*m.Buf[0][2] + v.W*m.Buf[0][3],
		Y: v.X*m.Buf[1][0] + v.Y*m.Buf[1][1] + v.Z*m.Buf[1][2] + v.W*m.Buf[1][3],
		Z: v.X*m.Buf[2][0] + v.Y*m.Buf[2][1] + v.Z*m.Buf[2][2] + v.W*m.Buf[2][3],
		W: v.X*m.Buf[3][0] + v.Y*m.Buf[3][1] + v.Z*m.Buf[3][2] + v.W*m.Buf[3][3],
	}
}
