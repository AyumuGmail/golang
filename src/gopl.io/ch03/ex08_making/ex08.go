package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
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
		width, height          = 256, 256
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img1 := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z1 := complex(x, y)
			img1.Set(px, py, mandelbrot64(complex64(z1)))
		}
	}
	out, _ := os.Create("img.png")
	out1, _ := os.Create("img1.png")
	png.Encode(out, img)
	png.Encode(out1, img1)
	//png.Encode(w, img)
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
func mandelbrot64(z complex64) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
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
func averrageColor(color1 color.RGBA, color2 color.RGBA, color3 color.RGBA, color4 color.RGBA) color.Color {
	var colorRave uint8 = uint8((int(color1.R) + int(color2.R) + int(color3.R) + int(color4.R)) / 4)
	var colorGave uint8 = uint8((int(color1.G) + int(color2.G) + int(color3.G) + int(color4.G)) / 4)
	var colorBave uint8 = uint8((int(color1.B) + int(color2.B) + int(color3.B) + int(color4.B)) / 4)

	if colorRave != color1.R {
		log.Printf("%d %d %d %d ave:%d\n", color1.R, color2.R, color3.R, color4.R, (int(color1.R)+int(color2.R)+int(color3.R)+int(color4.R))/4)
	}

	return color.RGBA{colorRave, colorGave, colorBave, 0xff}
}
