// Напишите программу общего назначения для конвертации единиц (аналогичную cf),
// которая читает числа из аргументов командной строки или из стандартного ввода (stdin),
// если аргументы не переданы, и преобразует каждое число в различные единицы:
// температуру в градусы Цельсия и Фаренгейта, длину в футы и метры, вес в фунты и килограммы и т.п.
package main

import (
	"TheGoProgrammingLanguage/Chapter02/tempconv"
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
)

type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

const (
	LengthConversionFactor float64 = 0.3048
	WeightConversionFactor float64 = 0.45359237
)

func (ft Foot) String() string     { return fmt.Sprintf("%g ft", ft) }
func (m Meter) String() string     { return fmt.Sprintf("%g m", m) }
func (lb Pound) String() string    { return fmt.Sprintf("%g lb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%g kg", kg) }

// FToM преобразует длину из футов в метры.
func FToM(ft Foot) Meter { return Meter(ft * Foot(LengthConversionFactor)) }

// MToF преобразует длину из метров в футы.
func MToF(m Meter) Foot { return Foot(m / Meter(LengthConversionFactor)) }

// PToK преобразует вес из фунтов в килограммы.
func PToK(lb Pound) Kilogram { return Kilogram(lb * Pound(WeightConversionFactor)) }

// KToP преобразует вес из килограммов в фунты.
func KToP(kg Kilogram) Pound { return Pound(kg / Kilogram(WeightConversionFactor)) }

func main() {
	if len(os.Args) == 1 {
		os.Stdout.WriteString("Введите значения (float64):\n")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			valueStr := scanner.Text()
			valueFloat64, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				logError(err.Error())
				continue
			}
			conversion(os.Stdout, valueFloat64)
		}
		if err := scanner.Err(); err != nil {
			logError(err.Error())
		}
		return
	}
	for _, arg := range os.Args[1:] {
		argFloat64, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			logError(err.Error())
			continue
		}
		conversion(os.Stdout, argFloat64)
	}
}

// conversion выводит в output конвертируемые значения.
func conversion(output io.Writer, value float64) {
	f := tempconv.Fahrenheit(value)
	c := tempconv.Celsius(value)
	ft := Foot(value)
	m := Meter(value)
	lb := Pound(value)
	kg := Kilogram(value)

	fmt.Fprintf(output, "%s = %s\n", f, tempconv.FToC(f))
	fmt.Fprintf(output, "%s = %s\n", c, tempconv.CToF(c))
	fmt.Fprintf(output, "%s = %s\n", ft, FToM(ft))
	fmt.Fprintf(output, "%s = %s\n", m, MToF(m))
	fmt.Fprintf(output, "%s = %s\n", lb, PToK(lb))
	fmt.Fprintf(output, "%s = %s\n", kg, KToP(kg))
	fmt.Fprintf(output, "\n")
}

// logError выводит ошибку, добавляя в начало файл, пакет и функцию в которой эта ошибка произошла.
func logError(msg string) {
	pc, file, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	fmt.Fprintf(os.Stderr, "%s:%s:%s\n", file, fn.Name(), msg)
}
