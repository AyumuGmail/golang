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

var palettes []string

func init() {
	//color.RGBA{0, 0xff, 0, 0xff}, //Green
	//color.White,
	//color.Black,
	//color.RGBA{0xff, 0, 0, 0xff}, //Red
	//color.RGBA{0, 0, 0xff, 0xff}, //Blue
	for i := 0; i <= int(0xff); i++ {
		palettes = append(palettes, fmt.Sprintf("#%02x%02x%02x", uint8(i), 0, 0xff-uint8(i)))

	}
}

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
	var points [cells][cells]XYpoint
	var maxZ float64
	var minZ float64 = 1000000

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, err1 := corner(i+1, j)
			bx, by, bz, err2 := corner(i, j)
			cx, cy, cz, err3 := corner(i, j+1)
			dx, dy, dz, err4 := corner(i+1, j+1)
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				//log.Println("Skipped!!")
				continue
			}

			var p XYpoint
			p.ax, p.ay, p.az, p.bx, p.by, p.bz, p.cx, p.cy, p.cz, p.dx, p.dy, p.dz =
				ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz
			points[i][j] = p
			maxZ = MaxZ(maxZ, az, bz, cz, dz)
			minZ = MinZ(minZ, az, bz, cz, dz)
		}
	}
	degreeValue := (maxZ - minZ) / 255

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			p := points[i][j]
			//degreeA := int((p.az + math.Abs(minZ)) / degreeValue)
			//degreeB := int((p.bz + math.Abs(minZ)) / degreeValue)
			//degreeC := int((p.cz + math.Abs(minZ)) / degreeValue)
			//Dの座標のみ描画の対象とする？
			degreeD := int((p.dz + math.Abs(minZ)) / degreeValue)
			//log.Printf("%d %d %d %d %s\n", degreeA, degreeB, degreeC, degreeD, palettes[degreeD])
			s = s + fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style=\"stroke:%s\"/>\n",
				p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, palettes[degreeD])
		}
	}
	s = s + fmt.Sprintf("</svg>\n")
	//log.Printf("maxZ:%f minZ:%f degreeValue:%f\n", maxZ, minZ, degreeValue)
	return s
}

type XYpoint struct {
	ax float64
	ay float64
	az float64
	bx float64
	by float64
	bz float64
	cx float64
	cy float64
	cz float64
	dx float64
	dy float64
	dz float64
}

func corner(i, j int) (float64, float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, err := f(x, y)

	sx := width/2 + (x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyrange - z*zscale

	return sx, sy, z, err
}

func f(x, y float64) (float64, error) {
	r := math.Hypot(x, y)
	ret := math.Sin(r) / r

	if math.IsNaN(ret) {
		//log.Printf("NaN? %f", r)
		return 0, errors.New("NaN")
	} else {
		return ret, nil
	}
}

func MaxZ(maxz float64, a float64, b float64, c float64, d float64) float64 {
	if maxz < a {
		maxz = a
	}

	if maxz < b {
		maxz = b
	}
	if maxz < c {
		maxz = c
	}
	if maxz < d {
		maxz = d
	}
	return maxz
}

func MinZ(minz float64, a float64, b float64, c float64, d float64) float64 {
	if minz > a {
		minz = a
	}

	if minz > b {
		minz = b
	}
	if minz > c {
		minz = c
	}
	if minz > d {
		minz = d
	}
	return minz
}
