// Modify the Lissajous server to read parameter values
// from the URL. For example, you might arrange it so
// that a URL like http://localhost:8000/?cycles=20 set
// the number of cycles to 20 instead of the default 5.
// Use strconv.Atoi function to convert the string
// parameter to an integer.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x0, 0xff, 0x0, 0xff}}

const (
	blackIndex = 0 //first color in palette
	greenIndex = 1 //next color in palette
)

type Param struct {
	Cycles  float64 // = 5     // number of complete x oscillator revolutions
	Res     float64 // = 0.001 // angular resolution
	Size    int     // = 100   // image canvas covers [-size..+size]
	Nframes int     // = 64    // number of animation frames
	Delay   int     // = 8     // delay between frames in 10ms units
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		rawQuery := r.URL.RawQuery
		values, err := url.ParseQuery(rawQuery)
		if err != nil {
			fmt.Fprintf(w, "while parsing query %s: %v", rawQuery, err)
			return
		}
		param := handleParams(values)
		lissajous(w, param)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleParams(values map[string][]string) Param {
	param := Param{Cycles: 5, Res: 0.001, Size: 100, Nframes: 64, Delay: 8}
	for key, value := range values {
		switch key {
		case "cycles":
			param.Cycles, _ = strconv.ParseFloat(value[0], 64)
		case "res":
			param.Res, _ = strconv.ParseFloat(value[0], 64)
		case "size":
			param.Size, _ = strconv.Atoi(value[0])
		case "nframes":
			param.Nframes, _ = strconv.Atoi(value[0])
		case "delay":
			param.Delay, _ = strconv.Atoi(value[0])
		}
	}
	return param
}

func lissajous(out io.Writer, param Param) {
	size := param.Size
	freq := rand.Float64() * 3.0 // relative frequency of y oscilator
	anim := gif.GIF{LoopCount: param.Nframes}
	phase := 0.0
	for i := 0; i < param.Nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < param.Cycles*2*math.Pi; t += param.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, param.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
