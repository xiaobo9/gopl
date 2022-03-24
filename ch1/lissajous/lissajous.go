package lissajous

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/xiaobo9/gopl/utils"
)

// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.White, color.RGBA{0x00, 0x00, 0xBB, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous() gif.GIF {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 65    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	return anim
}

func Lissajous() {
	fmt.Println(palette)
	fileName := fmt.Sprintf("lissajous_%v.gif", time.Now().Unix())
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixMicro())
	f, err := utils.CreateTempFile(fileName)
	log.Println(filepath.Abs(f.Name()))
	if err != nil {
		fmt.Println(err)
		return
	}
	anim := lissajous()
	err = gif.EncodeAll(f, &anim)
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
}

func LissajousOut(w io.Writer) {
	anim := lissajous()
	gif.EncodeAll(w, &anim)
}
