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
			c1 := mandelbrot(z1)
			c2 := mandelbrot(z2)
			c3 := mandelbrot(z3)
			c4 := mandelbrot(z4)
			img1.Set(px, py, averrageColor(c1, c2, c3, c4))
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

func averrageColor(color1 color.RGBA, color2 color.RGBA, color3 color.RGBA, color4 color.RGBA) color.Color {
	return color.RGBA{
		(color1.R + color2.R + color3.R + color4.R) / 4,
		(color1.G + color2.G + color3.G + color4.G) / 4,
		(color1.B + color2.B + color3.B + color4.B) / 4,
		0xff,
	}
}
