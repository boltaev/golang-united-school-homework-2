package square

import "math"

type sides int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sides) (square float64) {
	switch sidesNum {
	case 0:
		square = math.Pow(sideLen, 2)*math.Pi + SidesCircle
	case 4:
		square = sideLen * SidesSquare
	case 3:
		square = math.Pow(sideLen, 2) / 4 * math.Sqrt(SidesTriangle)
	default:
		square = 0
	}
	return
}
