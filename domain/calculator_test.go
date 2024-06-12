package calculator_test

import (
	calculator "devops/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("Testing with integers", func(t *testing.T) {
		cases := []struct {
			x, y, want int
		}{
			{1, 2, 3},
			{0, 0, 0},
			{-1, 1, 0},
			{-1, -2, -3},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Add(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})

	t.Run("Testing with float", func(t *testing.T) {
		cases := []struct {
			x, y, want float64
		}{
			{1., 2., 3.},
			{0., 0., 0.},
			{-1., 1., 0.},
			{-1., -2., -3.},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Add(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("Testing with integers", func(t *testing.T) {
		cases := []struct {
			x, y, want int
		}{
			{1, 2, -1},
			{0, 0, 0},
			{-1, 1, -2},
			{-1, -2, 1},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Sub(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})

	t.Run("Testing with float", func(t *testing.T) {
		cases := []struct {
			x, y, want float64
		}{
			{1., 2., -1.},
			{0., 0., 0.},
			{-1., 1., -2.},
			{-1., -2., 1.},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Sub(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})
}

func TestMul(t *testing.T) {
	t.Run("Testing with integers", func(t *testing.T) {
		cases := []struct {
			x, y, want int
		}{
			{1, 2, 2},
			{0, 0, 0},
			{-1, 1, -1},
			{-1, -2, 2},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Mul(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})

	t.Run("Testing with float", func(t *testing.T) {
		cases := []struct {
			x, y, want float64
		}{
			{1., 2., 2.},
			{0., 0., 0.},
			{-1., 1., -1.},
			{-1., -2., 2.},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Mul(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})
}

func TestDiv(t *testing.T) {
	t.Run("Testing with integers", func(t *testing.T) {
		cases := []struct {
			x, y, want int
		}{
			{1, 2, 0},
			{-1, 1, -1},
			{-1, -2, 0},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Div(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})

	t.Run("Testing with float", func(t *testing.T) {
		cases := []struct {
			x, y, want float64
		}{
			{1., 2., 1. / 2.},
			{-1., 1., -1.},
			{-1., -2., 1. / 2.},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Div(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})
}

func TestPow(t *testing.T) {
	t.Run("Testing with integers", func(t *testing.T) {
		cases := []struct {
			x, y, want int
		}{
			{1, 2, 1},
			{-1, 1, -1},
			{2, 2, 4},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Pow(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})

	t.Run("Testing with float", func(t *testing.T) {
		cases := []struct {
			x, y, want float64
		}{
			{1., 2., 1.},
			{-1., 1., -1.},
			{2., 2., 4.},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("Case_%d", i), func(t *testing.T) {
				res := calculator.Pow(c.x, c.y)
				assert.Equal(t, res, c.want)
			})
		}
	})
}
