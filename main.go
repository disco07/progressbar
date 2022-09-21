package main

import (
	"fmt"
	"time"
)

type Bar struct {
	percent int64
	current int64
	total   int64
	rate    string
	graph   string
}

func NewOption(end int64) *Bar {
	current := int64(0)
	total := end
	graph := "#"
	percent := getPercent(current, total)

	return &Bar{
		percent: percent,
		current: current,
		total:   total,
		graph:   graph,
	}
}

func getPercent(current, total int64) int64 {
	return int64((float32(current) / float32(total)) * 100)
}

func (bar *Bar) PlayBar(current int64) {
	bar.current = current
	last := bar.percent
	bar.percent = getPercent(current, bar.total)
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.current, bar.total)
}

func Default(end int64) *Bar {
	return NewOption(end)
}

func main() {
	bar := Default(50)
	for i := 0; i <= 50; i++ {
		time.Sleep(50 * time.Millisecond)
		bar.PlayBar(int64(i))
	}
}
