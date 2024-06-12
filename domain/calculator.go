package calculator

import "golang.org/x/exp/constraints"

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
