package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var (
	defaultWidth, defaultHeight = 600, 320    // размер холста в пикселях
	cells                       = 100         // количество ячеек сетки
	xyrange                     = 30.0        // диапазон осей (-xyrange..+xyrange)
	angle                       = math.Pi / 6 // угол осей x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func handler(w http.ResponseWriter, r *http.Request) {
	width := defaultWidth
	height := defaultHeight
	stroke := "grey"
	fill := "white"

	queries := r.URL.Query()
	if valueString := queries.Get("width"); valueString != "" {
		valueInt, err := strconv.Atoi(valueString)
		if err == nil {
			width = valueInt
		} else {
			fmt.Printf("Incorrect query width: %v", err)
		}
	}
	if valueString := queries.Get("height"); valueString != "" {
		valueInt, err := strconv.Atoi(valueString)
		if err == nil {
			height = valueInt
		} else {
			fmt.Printf("Incorrect query height: %v", err)
		}
	}
	if value := queries.Get("stroke"); value != "" {
		stroke = value
	}
	if value := queries.Get("fill"); value != "" {
		fill = value
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(surface(width, height, stroke, fill)))

}

func surface(width, height int, stroke, fill string) string {
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4

	var svg strings.Builder
	fmt.Fprintf(&svg, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", stroke, fill, width, height)
	for i := range cells {
		for j := range cells {
			ax, ay := corner(i+1, j, width, height, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyscale, zscale)
			fmt.Fprintf(&svg, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(&svg, "</svg>")
	return svg.String()
}

func corner(i, j, width, height int, xyscale, zscale float64) (float64, float64) {
	// Находим точку (x, y) в углу ячейки (i, j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Вычисляем высоту поверхность z.
	z := f(x, y)

	// Изометрически проецируем (x, y, z) на 2D SVG-холст (sx, sy)
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // расстояние от (0, 0)
	if r == 0 {
		return 1
	}
	return math.Sin(r) / r
}
