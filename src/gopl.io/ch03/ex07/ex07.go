package main

import (
	"image"
	"image/color"
	"image/png"
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
		xmin, ymin, xmax, ymax = -1, -1, +1, +1
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	out, _ := os.Create("img.png")
	png.Encode(out, img)
}

const (
	RGBRed, RGBGreen, RGBBlue, RGBAlpha = 150, 15, 200, 0xff
)

func newton(z complex128) color.RGBA {
	const iterations = 255
	const contrast = 15

	var v complex128
	var v2 complex128
	v = z
	for n := uint8(0); n < iterations; n++ {
		v2 = (v - (v * v * v * v / (4 * v * v * v)))
		//if math.Abs(v2-v) < 0.00001 {
		//
		//	}
		//ニュートン法と収束の学習
		v = v2
	}
	return color.RGBA{0, 0, 0, RGBAlpha} //color.Black
}
