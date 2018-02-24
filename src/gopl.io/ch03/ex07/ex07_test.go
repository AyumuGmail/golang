package main

import (
	"fmt"
	"testing"
)

func TestNewton(t *testing.T) {
	var tests = []struct {
		x float64
		y float64
	}{
		{1, 1}, {-1, 1}, {0.12, 0.3}, {-0.23, -0.12}, {0, 0},
	}
	for _, test := range tests {
		results := newton(complex(test.x, test.y))
		fmt.Printf("%v\n", results)
	}
}

func TestNewtonLoop(t *testing.T) {
	const (
		xmin, ymin, xmax, ymax = -1, -1, 1, +1
		width, height          = 8, 8
	)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//img.Set(px, py, newton(z))
			newton(z)
		}
	}
}
