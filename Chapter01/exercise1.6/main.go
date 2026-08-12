// Измените программу lissajous так, чтобы она генерировала изображения разных цветов,
// добавляя в палитру palette больше значений, а затем выводя их путем изменения
// третьего аргумента функции SetColorlndex некоторым нетривиальным способом.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
	"os"
)

var palette = color.Palette{
	color.White,
	color.RGBA{255, 0, 0, 255},   // Красный
	color.RGBA{255, 177, 0, 255}, // Оранжевый
	color.RGBA{255, 255, 0, 255}, // Желтый
	color.RGBA{0, 255, 0, 255},   // Зеленый
	color.RGBA{0, 255, 255, 255}, // Голубой
	color.RGBA{0, 0, 255, 255},   // Синий
	color.RGBA{128, 0, 128, 255}, // Фиолетовый
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 25     // Количество полных колебаний x
		res     = 0.0001 // Угловое разрешение
		size    = 100    // Канва изображения охватывает [size..+size]
		nframes = 64     // Количество кадров анимации
		delay   = 8      // Задержка между кадрами (единица - 10мс)
	)
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(int(t)%(len(palette)-1) + 1)
			//img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(palette.Index(generateRandomRGB())))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}

func generateRandomRGB() color.Color {
	R := uint8(rand.IntN(256))
	G := uint8(rand.IntN(256))
	B := uint8(rand.IntN(256))
	return color.RGBA{R, G, B, 255}
}
