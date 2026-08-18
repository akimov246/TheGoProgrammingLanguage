// Измените сервер с фигурами Лиссажу так, чтобы значения
// параметров считывались из URL. Например, URL вида http://localhost:8000/?cycles=20
// устанавливает количество циклов равным 20 вместо значения по умол­чанию, равного 5.
// Используйте функцию strconv.Atoi для преобразования строкового параметра в целое число.
// Просмотреть документацию по данной функции можно с помощью команды go doc strconv.Atoi.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // Первый цвет палитры
	blackIndex = 1 // Следующий цвет палитры
)

type lissajousParams struct {
	cycles  int     // Количество полных колебаний x
	res     float64 // Угловое разрешение
	size    int     // Канва изображения охватывает [size..+size]
	nframes int     // Количество кадров анимации
	delay   int     // Задержка между кадрами (единица - 10мс)
}

// Lissajous генерирует анимированный GIF из случайных фигур Лиссажу.
func lissajous(out io.Writer, params lissajousParams) {
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: params.nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < params.nframes; i++ {
		rect := image.Rect(0, 0, 2*params.size+1, 2*params.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(params.cycles)*2*math.Pi; t += params.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(params.size+int(x*float64(params.size)+0.5), params.size+int(y*float64(params.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := lissajousParams{
		cycles:  5,
		res:     0.0001,
		size:    100,
		nframes: 64,
		delay:   8,
	}

	queryParams := r.URL.Query()

	allowedQueryParams := map[string]bool{
		"cycles":  true, // Количество полных колебаний x
		"res":     true, // Угловое разрешение
		"size":    true, // Канва изображения охватывает [size..+size]
		"nframes": true, // Количество кадров анимации
		"delay":   true, // Задержка между кадрами (единица - 10мс)
	}

	// Нет лишних query-параметров.
	for queryParam, value := range queryParams {
		if !allowedQueryParams[queryParam] {
			http.Error(w, fmt.Sprintf("unexpected query parameter %q", queryParam), http.StatusBadRequest)
			return
		}
		if len(value) > 1 {
			http.Error(w, fmt.Sprintf("the query parameter value %q must be specified exactly once", queryParam), http.StatusBadRequest)
			return
		}
	}

	if value := queryParams.Get("cycles"); value != "" {
		valueInt, err := strconv.Atoi(value)
		if err != nil || valueInt < 1 {
			http.Error(w, fmt.Sprintf("the query-parameter value \"cycles\" must be of type %T and greater than zero", params.cycles), http.StatusBadRequest)
			return
		}
		params.cycles = valueInt
	}
	if value := queryParams.Get("size"); value != "" {
		valueInt, err := strconv.Atoi(value)
		if err != nil || valueInt < 1 {
			http.Error(w, fmt.Sprintf("the query-parameter value \"size\" must be of type %T and greater than zero", params.size), http.StatusBadRequest)
			return
		}
		params.size = valueInt
	}
	if value := queryParams.Get("nframes"); value != "" {
		valueInt, err := strconv.Atoi(value)
		if err != nil || valueInt < 1 {
			http.Error(w, fmt.Sprintf("the query-parameter value \"nframes\" must be of type %T and greater than zero", params.nframes), http.StatusBadRequest)
			return
		}
		params.nframes = valueInt
	}
	if value := queryParams.Get("delay"); value != "" {
		valueInt, err := strconv.Atoi(value)
		if err != nil || valueInt < 1 {
			http.Error(w, fmt.Sprintf("the query-parameter value \"delay\" must be of type %T and greater than zero", params.delay), http.StatusBadRequest)
			return
		}
		params.delay = valueInt
	}
	if value := queryParams.Get("res"); value != "" {
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil || valueFloat64 <= 0 {
			http.Error(w, fmt.Sprintf("the query-parameter value \"res\" must be of type %T and greater than zero", params.res), http.StatusBadRequest)
			return
		}
		params.res = valueFloat64
	}
	lissajous(w, params)
}
