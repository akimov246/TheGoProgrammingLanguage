// Tempconv выполняет вычисления температур по Цельюсию (Celsium) и по Фаренгейту (Fahrenheit).
package tempconv

type Celsium float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsium = -273.15
	FreezingC     Celsium = 0
	BoilingC      Celsium = 100
)

func CToF(c Celsium) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsium {
	return Celsium((f - 32) * 5 / 9)
}
