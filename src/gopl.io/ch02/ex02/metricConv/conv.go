package metricConv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//CtoK is xxx
func CToK(c Celsius) Kelvin {
	return (Kelvin)(c + AbsoluteZeroC)
}

func Pd2kg(p Pond) KiloGram {
	return KiloGram(p / 2.2046)
}

func Kg2pd(k KiloGram) Pond {
	return Pond(k * 2.2046)
}

func M2ft(m Meter) Feet {
	return Feet(m * 3.2808)
}

func Ft2m(f Feet) Meter {
	return Meter(f / 3.2808)
}
