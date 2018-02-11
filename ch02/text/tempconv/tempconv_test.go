package tempconv

import (
	"fmt"
	"math"
	"testing"
)

func TestExec(t *testing.T) {
	var tests = []struct {
		c Celsius
		k Kelvin
	}{
		{0, -273.15},
		{100, -173.1500000},
	}
	for _, test := range tests {
		k := CToK(test.c)
		k = Kelvin(Round(float64(k), 2))
		if k != test.k {
			t.Errorf("k error:expected:%v actual:%v", test.k, k)
		} else {
			fmt.Println("Success")
		}
	}
}

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
