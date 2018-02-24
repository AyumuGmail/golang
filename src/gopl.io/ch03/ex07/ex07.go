package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
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
		xmin, ymin, xmax, ymax = -1, -1, 1, +1
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
	RGBRed, RGBGreen, RGBBlue, RGBAlpha = 200, 200, 200, 0xff
)

func newton(z complex128) color.RGBA {
	const iterations = 255
	const contrast = 15
	const diffendVal = 0.000001

	var v complex128
	var v2 complex128
	v = z
	for n := uint8(0); n < iterations; n++ {
		v2 = (v - ((v*v*v*v - 1) / (4 * v * v * v)))
		//if math.Abs(v2-v) < 0.00001 {
		//
		//	}
		//ニュートン法と収束の学習
		if cmplx.IsNaN(v2) {
			fmt.Printf("****** Nan")
			return color.RGBA{0, 0, 0, RGBAlpha} //color.Black
		}

		diffX := real(v2) - real(v)
		diffY := imag(v2) - imag(v)
		if math.Abs(diffX) < diffendVal && math.Abs(diffY) < diffendVal {
			if diffX > 0 && diffY > 0 {
				fmt.Printf("*両方正。%d回目\n", n)
				return color.RGBA{RGBRed + contrast*n,
					RGBGreen,
					RGBBlue,
					RGBAlpha}
			} else if diffX < 0 && diffY > 0 {
				fmt.Printf("**Xが負の数字。%d回目\n", n)
				return color.RGBA{RGBRed,
					RGBGreen + contrast*n,
					RGBBlue,
					RGBAlpha}
			} else if diffX > 0 && diffY < 0 {
				fmt.Printf("*** Yが負の数字。%d回目\n", n)
				return color.RGBA{RGBRed,
					RGBGreen,
					RGBBlue + contrast*n,
					RGBAlpha}
			} else if diffX < 0 && diffY < 0 {
				fmt.Printf("**** XもYも負の数字。%d回目\n", n)
				return color.RGBA{contrast * n,
					contrast * n,
					contrast * n,
					RGBAlpha}
			}
		}
		//fmt.Printf("%d回目:%f %f real:%f imag:%f %f\n", n, v2, cmplx.Abs(v2-v), diffX, diffY, complex(0, -0.0000000001))
		v = v2
	}
	return color.RGBA{0, 0, 0, RGBAlpha} //color.Black
}
