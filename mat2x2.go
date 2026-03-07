package kmath

type Mat2x2 struct {
	Buf [2][2]float64
}

func MakeMat2x2(b [2][2]float64) Mat2x2 {
	return Mat2x2{
		Buf: b,
	}
}

func (m2 *Mat2x2) Add(other *Mat2x2) Mat2x2 {
	var newM Mat2x2

	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			newM.Buf[y][x] = m2.Buf[y][x] + other.Buf[y][x]
		}
	}

	return newM
}

func (m2 *Mat2x2) Sub(other *Mat2x2) Mat2x2 {
	var newM Mat2x2

	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			newM.Buf[y][x] = m2.Buf[y][x] - other.Buf[y][x]
		}
	}

	return newM
}

func (m2 *Mat2x2) Mul(other *Mat2x2) Mat2x2 {
	var newM Mat2x2

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			var sum float64
			for k := 0; k < 2; k++ {
				sum += m2.Buf[y][k] * other.Buf[k][x]
			}
			newM.Buf[y][x] = sum
		}
	}

	return newM
}
