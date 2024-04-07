package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)
var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func main() {
	lissajous()
}

func lissajous() {
	//lissajous1(os.Stdout)
	lissajous2(os.Stdout)
}

func lissajous1(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 128
		nframes = 256
		delay   = 4
	)
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	frequency := r.Float64() * 3.0
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rectangle := image.Rect(0, 0, 2*size+1, 2*size+1)
		image := image.NewPaletted(rectangle, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			image.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, image)
	}
	gif.EncodeAll(out, &animation)
}

func lissajous2(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 128
		nframes = 256
		delay   = 4
	)
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	frequency := r.Float64() * 3.0
	//fmt.Println(r.Float64());
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rectangle := image.Rect(0, 0, 2*size+1, 2*size+1)
		image := image.NewPaletted(rectangle, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			image.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, image)
	}
	gif.EncodeAll(out, &animation)
}
