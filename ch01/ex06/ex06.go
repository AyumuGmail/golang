package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{0, 0, 0, 1},
	color.RGBA{0, 255, 0, 255}, //Green
	color.RGBA{255, 0, 0, 255}, //Red
	color.RGBA{0, 0, 255, 255}, //Blue
}

const (
	backgroundIndex = 0
	lineColorIndex  = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	colorsize := len(palette)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		//img := image.NewRGBA(rect)
		j := 0
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			j++
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := (uint8)(j % colorsize)
			//			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), lineColorIndex)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
