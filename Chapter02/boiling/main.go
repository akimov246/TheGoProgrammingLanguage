// Boiling выводит температуру кипения воды.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Теспература кипения = %g°F или %g°C\n", f, c)
}
