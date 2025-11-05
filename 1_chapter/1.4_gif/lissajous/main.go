// Генерирует анимированные GIF из случайных фигур Лиссажу
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var palette = []color.Color{
	color.Black,                        // 0 - черный фон
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // 1 - зеленый
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // 2 - красный
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // 3 - синий
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // 4 - желтый
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // 5 - пурпурный
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // 6 - голубой
	color.RGBA{0xFF, 0x80, 0x00, 0xFF}, // 7 - оранжевый
	color.RGBA{0x80, 0x00, 0xFF, 0xFF}, // 8 - фиолетовый
}

const (
	blackIndex = 0 // Следующий щвет палитры
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	http.ListenAndServe(":8080", nil)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задерка между кадрами (единица 10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота калебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		colorIndex := selectColorIndex(i, nframes)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Игнорируем ошибки
}

func selectColorIndex(frame, totalFrames int) uint8 {
	return phaseBasedColor(frame)
}

func phaseBasedColor(frame int) uint8 {
	phase := float64(frame) * 0.1
	// Используем синус для плавного перехода между цветами
	value := math.Sin(phase)
	// Преобразуем в диапазон 1-8
	colorIndex := uint8((value+1)/2*7 + 1)
	return colorIndex
}
