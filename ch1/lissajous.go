// Lissajous gen GIF animations of random Lissajous figures
// Run: 
//    go build lissajous.go
//    ./lissajous > out.gif

package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
)

var palette = []color.Color{color.White, color.Black}

const (
    WhiteIndex = 0
    BlackIndex = 1
)

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
       cycle = 5  // number of complete x osclillator revolutions
       res = 0.0001  // angular resolution
       size = 100  // image convas covers
       nframes = 64  // number of amimation frames
       delay = 8  // delay between frames in 10ms unit.
    )

    freq := rand.Float64() * 3.0 //relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)

        for t := 0.0; t < cycle * 2 * math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            img.SetColorIndex(
                size + int(x * size + 0.5),
                size + int(y * size + 0.5),
                BlackIndex)
        }

        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)

    }
    // Note: ignoring encoding errors
    gif.EncodeAll(out, &anim)

}
