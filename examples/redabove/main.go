package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
)

func main() {
	series := []float64{1, 2, 3, 4, 5, 4, 3, 2, 1}
	graph := asciigraph.Plot(series, asciigraph.Height(10),
		asciigraph.Width(50),
		asciigraph.Caption("A simple line graph"),
		asciigraph.ColorAbove(asciigraph.Red, 4.0),
		asciigraph.ColorBelow(asciigraph.DarkGreen, 2.0),
	)
	fmt.Println(graph)
}
