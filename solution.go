package square

import "math"

type sides int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum sides) (square float64) {
	switch sidesNum {
	case 0:
		square = math.Pow(sideLen, 2)*math.Pi + SidesCircle
	case 4:
		square = sideLen * SidesSquare
	case 3:
		square = math.Pow(sideLen, 2) / 4 * math.Sqrt(SidesTriangle)
	}
	return
}
