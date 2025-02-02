package square

import "math"

type sides int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

// CalcSquare formulas changed a bit(except case 3) cause of using all constants
func CalcSquare(sideLen float64, sidesNum sides) (square float64) {
	switch sidesNum {
	case 0:
		square = math.Pow(sideLen, 2)*math.Pi + SidesCircle
	case 4:
		square = math.Pow(sideLen, SidesSquare/2)
	case 3:
		square = math.Pow(sideLen, 2) / 4 * math.Sqrt(SidesTriangle)
	}
	return
}
