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
    "sync"
)

var (
    mu    sync.Mutex
    count int
)

var palette2 = []color.Color{color.White, color.Black, color.RGBA{G: 255, A: 1}}

const (
    whiteIndex2 = 0 // first color in palette
    blackIndex2 = 1 // next color in palette
    greenIndex2 = 2 // next color in palette
)


func main() {
    http.HandleFunc("/", echo)
    http.HandleFunc("/count", counter)
    log.Fatalln(http.ListenAndServe(":8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "count %d\n", count)
    fmt.Fprintf(w, "%s\t%s\t%s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
    mu.Unlock()
}

// handler echoes the Path component of the requested URL.
func echo(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    lissajous2(w)
    mu.Unlock()
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func lissajous2(out io.Writer) {
    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 150   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette2)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex2)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    err := gif.EncodeAll(out, &anim)
    if err != nil {
        return
    }
}
