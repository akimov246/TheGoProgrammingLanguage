package tempconv

// CToF преобразует температуру по Цельсию в Фаренгейт.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC преобразует температуру по Фаренгейту в Цельсий.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
