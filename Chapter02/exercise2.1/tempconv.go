// Добавьте типы, константы и функции в tempconv для работы с температурой по шкале Кельвина,
// где ноль Кельвина равен −273.15°C, а разница в 1K имеет ту же величину, что и 1°C.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// CToF преобразует температуру по Цельсию в Фаренгейт.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK преобразует температуру по Цельсию в Кельвины.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FToC преобразует температуру по Фаренгейту в Цельсий.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK преобразует температуру по Фаренгейту в Кельвины.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToC преобразует температуру по Кельвину в Цельсий.
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

// KToF преобразует температуру по Кельвину в Фаренгейты.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
