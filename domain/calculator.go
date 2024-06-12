package calculator

import (
	"golang.org/x/exp/constraints"
	"math"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](x, y T) T {
	return x + y
}

func Sub[T Number](x, y T) T {
	return x - y
}

func Mul[T Number](x, y T) T {
	return x * y
}

func Div[T Number](x, y T) T {
	return x / y
}

func Pow[T Number](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}
