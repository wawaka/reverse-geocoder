package main

import (
    "fmt"
)

func main() {
    // float32 x = 0;
    // var mpd float32
    var dpm float32
    dpm = float32(1/meters_per_degree_lon(0))
    var x, y float32
    x = 105
    y = x
    y += dpm
    y += dpm
    fmt.Printf("x=%.20f\n", x)
    fmt.Printf("y=%.20f\n", y)
    fmt.Printf("d=%.20f\n", (y-x)/dpm)
}
