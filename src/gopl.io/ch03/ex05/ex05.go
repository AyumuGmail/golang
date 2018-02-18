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
	img1 := image.NewRGBA(image.Rect(0, 0, width, height))
	var sampllingX float64 = (float64(1) / width) * ((xmax - xmin) / 2)
	var sampllingY float64 = (float64(1) / height) * ((ymax - ymin) / 2)

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
			z2 := complex(x+sampllingX, y)
			z3 := complex(x, y+sampllingY)
			z4 := complex(x+sampllingX, y+sampllingY)
			img1.Set(px, py, mandelbrotSuperSamplling(z1, z2, z3, z4))
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

func mandelbrot(z complex128) color.Color {
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
	return color.Black
}

func mandelbrotSuperSamplling(z1 complex128, z2 complex128, z3 complex128, z4 complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex128((z1+z2+z3+z4)/4)
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
	return color.Black
}
