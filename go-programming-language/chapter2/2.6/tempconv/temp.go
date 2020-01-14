package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -237.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin(273.15 + (f-32)*5/9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(273.15 + c)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
