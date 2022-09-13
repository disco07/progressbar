package main

import (
	"fmt"
)

type Bar struct {
	percent int64
	current int64
	total   int64
	rate    string
	graph   string
}

func (bar *Bar) NewOption(start, end int64) {
	bar.current = start
	bar.total = end
	if bar.graph == "" {
		bar.graph = "#"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 1 {
		bar.rate += bar.graph
	}
}

func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.current) / float32(bar.total)) * 100)
}

func (bar *Bar) PlayBar(current int64) {
	bar.current = current
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.current, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println("\nFinished")
}

func ProgressBar(end int64) {
	var bar Bar
	bar.NewOption(0, end)
	for i := 0; i <= int(end); i++ {
		bar.PlayBar(int64(i))
	}
	bar.Finish()
}

func main() {
	ProgressBar(100)
}
