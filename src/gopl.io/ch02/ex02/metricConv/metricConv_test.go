package metricConv

import (
	"math"
	"testing"
	//	"gopl.io/ch02/ex02/metricConv"
)

func TestExec(t *testing.T) {
	var tests = []struct {
		c  Celsius
		k  Kelvin
		f  Feet
		m  Meter
		p  Pond
		kg KiloGram
	}{
		{0, -273.15, 100, 30.48, 100, 45.36},
	}
	for _, test := range tests {
		p := Pond(Round(float64(Kg2pd(test.kg)), 2))
		kg := KiloGram(Round(float64(Pd2kg(test.p)), 2))
		f := Feet(Round(float64(M2ft(test.m)), 2))
		m := Meter(Round(float64(Ft2m(test.f)), 2))

		if p != test.p {
			t.Errorf("expected:%v actual:%v", test.p, p)
		}
		if kg != test.kg {
			t.Errorf("expected:%v actual:%v", test.kg, kg)
		}
		if f != test.f {
			t.Errorf("expected:%v actual:%v", test.f, f)
		}
		if m != test.m {
			t.Errorf("expected:%v actual:%v", test.m, m)
		}
	}
}

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
