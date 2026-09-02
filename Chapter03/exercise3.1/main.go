package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // размер холста в пикселях
	cells         = 100                 // количество ячеек сетки
	xyrange       = 30.0                // диапазон осей (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // пикселей на единицу x или y
	zscale        = height * 0.4        // пикселей на единицу z
	angle         = math.Pi / 6         // угол осей x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, Aok := corner(i+1, j)
			bx, by, Bok := corner(i, j)
			cx, cy, Cok := corner(i, j+1)
			dx, dy, Dok := corner(i+1, j+1)
			if !Aok || !Bok || !Cok || !Dok {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Находим точку (x, y) в углу ячейки (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Вычисляем высоту поверхность z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}

	// Изометрически проецируем (x, y, z) на 2D SVG-холст (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // расстояние от (0, 0)
	res := math.Sin(r) / r
	if math.IsInf(res, 0) || math.IsNaN(res) {
		return 0, false
	}
	return res, true
}
