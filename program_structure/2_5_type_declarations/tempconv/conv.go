package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

// ex_2_1
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius ((f - 32) * 5 / 9)
}

// ex_2_1
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}
