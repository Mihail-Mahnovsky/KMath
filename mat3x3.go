package kmath

type Mat3x3 struct {
	Buf [3][3]float64
}

func MakeMat3x3(b [3][3]float64) Mat3x3 {
	return Mat3x3{
		Buf: b,
	}
}

func (m3 *Mat3x3) Add(other *Mat3x3) Mat3x3 {
	var newM Mat3x3

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			newM.Buf[y][x] = m3.Buf[y][x] + other.Buf[y][x]
		}
	}

	return newM
}

func (m3 *Mat3x3) Sub(other *Mat3x3) Mat3x3 {
	var newM Mat3x3

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			newM.Buf[y][x] = m3.Buf[y][x] - other.Buf[y][x]
		}
	}

	return newM
}

func (m3 *Mat3x3) Mul(other *Mat3x3) Mat3x3 {
	var newM Mat3x3

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			var sum float64
			for k := 0; k < 3; k++ {
				sum += m3.Buf[y][k] * other.Buf[k][x]
			}
			newM.Buf[y][x] = sum
		}
	}

	return newM
}

func (m *Mat3x3) MulVec(v Vec3) Vec3 {
	return Vec3{
		X: v.X*m.Buf[0][0] + v.Y*m.Buf[0][1] + v.Z*m.Buf[0][2],
		Y: v.X*m.Buf[1][0] + v.Y*m.Buf[1][1] + v.Z*m.Buf[1][2],
		Z: v.X*m.Buf[2][0] + v.Y*m.Buf[2][1] + v.Z*m.Buf[2][2],
	}
}
