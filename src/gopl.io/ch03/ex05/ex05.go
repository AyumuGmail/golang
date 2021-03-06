package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	mainProc(w)
	//})
	mainProc()
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func mainProc() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	out, _ := os.Create("img.png")
	png.Encode(out, img)
}

const (
	RGBRed, RGBGreen, RGBBlue, RGBAlpha = 150, 15, 200, 0xff
)

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			//フルカラーでも同じ色っていうのが正解?
			//return color.RGBA{255 - contrast*n, 255 - contrast*n, 255 - contrast*n, 0xff}
			return color.RGBA{RGBRed - contrast*n,
				RGBGreen - contrast*n,
				RGBBlue - contrast*n,
				RGBAlpha}
		}
	}
	return color.RGBA{0, 0, 0, RGBAlpha} //color.Black
}
