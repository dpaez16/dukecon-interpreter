package evaluator

import "math"

// https://voidpoint.io/terminx/eduke32/-/blob/master/source/duke3d/src/gameexec.cpp?ref_type=heads

func Sin(x int64) int64 {
	angle := buildEngineAngleToRadians(x)
	value := math.Sin(angle) * math.Pow(2, 14)

	if value < 0 {
		return int64(math.Ceil(value))
	} else {
		return int64(math.Floor(value))
	}
}

func Cos(x int64) int64 {
	return Sin(x + 512)
}

func Sqrt(x int64) int64 {
	value := math.Floor(math.Sqrt(float64(x)))
	return int64(value)
}

func CalcHypotenuse(x int64, y int64) int64 {
	hypot := x*x + y*y
	return Sqrt(hypot)
}

func GetAngle(x int64, y int64) int64 {
	if x == 0 && y == 0 {
		return 0
	}

	angle := math.Atan2(float64(y), float64(y))
	return radiansToBuildEngineAngle(angle)
}

func GetIncAngle(angle1 int64, angle2 int64) int64 {
	x := angle1 & 2047
	y := angle2 & 2047

	if Abs(x-y) < 1024 {
		return y - x
	}

	if y > 1024 {
		y -= 2048
	}

	if x > 1024 {
		x -= 2048
	}

	return y - x
}

func ShiftVarL(x int64, n int) int64 {
	return x << n
}

func ShiftVarR(x int64, n int) int64 {
	return x >> n
}

func MulScale(x int64, y int64, n int) int64 {
	return ShiftVarR(x*y, n)
}

func DivScale(x int64, d int64, n int) int64 {
	return Div(ShiftVarL(x, n), d)
}

func ScaleVar(x int64, y int64, d int64) int64 {
	return Div(x*y, d)
}

func Mul(x int64, y int64) int64 {
	return x * y
}

func Div(x int64, y int64) int64 {
	return x / y
}

func Divr(x int64, y int64) int64 {
	value := float64(x) / float64(y)
	return int64(math.RoundToEven(value))
}

func Divru(x int64, y int64) int64 {
	value := float64(x) / float64(y)

	if value > 0 {
		return int64(math.Ceil(value))
	} else {
		return int64(math.Floor(value))
	}
}

func Divrd(x int64, y int64) int64 {
	return Div(x, y)
}

func Abs(x int64) int64 {
	return int64(math.Abs(float64(x)))
}

func Inv(x int64) int64 {
	return -x
}

func Clamp(x int64, a int64, b int64) int64 {
	if x < a {
		return a
	}

	if x > b {
		return b
	}

	return x
}

func buildEngineAngleToRadians(x int64) float64 {
	x = x & 2047
	return 2 * math.Pi * (float64(x) / 2048)
}

func radiansToBuildEngineAngle(x float64) int64 {
	value := (2048 * x) / (2 * math.Pi)
	return int64(value)
}
