package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cyclesStr := r.FormValue("cycles")
	if cyclesStr == "" {
		cyclesStr = "20"
	}
	cycles, _ := strconv.Atoi(cyclesStr)
	lissajous(w, cycles)
}

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

func lissajous(out io.Writer, paramCycles int) {
	const (
		//	cycles  = cycles
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	cycles := float64(paramCycles)
	log.Print(cycles)

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
