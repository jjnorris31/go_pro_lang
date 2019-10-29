package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 0
)

// sobreescribe el método String()
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// sobreescribe el método String()
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// sobreescribe el método String()
func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k);
}