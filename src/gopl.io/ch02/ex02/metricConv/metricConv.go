package metricConv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meter float64
type Pond float64
type KiloGram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}
func (kg KiloGram) String() string {
	return fmt.Sprintf("%g Kg", kg)
}
func (p Pond) String() string {
	return fmt.Sprintf("%g lb", p)
}
func (m Meter) String() string {
	return fmt.Sprintf("%g M", m)
}
func (f Feet) String() string {
	return fmt.Sprintf("%g ft", f)
}
