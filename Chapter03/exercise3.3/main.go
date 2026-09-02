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
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			colorRBG := getRBG(az, bz, cz, dz)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, colorRBG)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Находим точку (x, y) в углу ячейки (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Вычисляем высоту поверхность z.
	z := f(x, y)

	// Изометрически проецируем (x, y, z) на 2D SVG-холст (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // расстояние от (0, 0)
	return math.Sin(r) / r
}

func getRBG(az float64, bz float64, cz float64, dz float64) string {
	z := (az + bz + cz + dz) / 4
	minZ := -0.2
	maxZ := 0.2
	t := (z - minZ) / (maxZ - minZ)
	r := int(255 * t)
	g := 0
	b := int(255 * (1 - t))
	return fmt.Sprintf("rgb(%d, %d, %d)", int(r), int(g), int(b))
}
