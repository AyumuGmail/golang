package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var tan30 = math.Tan(angle)
var flr = math.Floor(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		s := Svg()
		fmt.Fprintf(w, "%s", s)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	//s := Svg()
	//fmt.Print(s)
}

func Svg() string {
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey;fill:white;stroke-width:0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err1 := corner(i+1, j)
			bx, by, err2 := corner(i, j)
			cx, cy, err3 := corner(i, j+1)
			dx, dy, err4 := corner(i+1, j+1)
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				log.Println("Skipped!!")
				continue
			}
			s = s + fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	s = s + fmt.Sprintf("</svg>\n")
	return s
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, err := f(x, y)

	sx := width/2 + (x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyrange - z*zscale

	return sx, sy, err
}

func f(x, y float64) (float64, error) {
	//r := math.Tan(y)
	r := math.Hypot(x, y)
	ret := math.Sin(r) / r

	if math.IsNaN(ret) {
		log.Printf("NaN? %f", r)
		return 0, errors.New("NaN")
	} else {
		return ret, nil
	}
}
