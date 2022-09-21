package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

type Bar struct {
	percent int64
	current int64
	rate    string
	graph   string
	config  config
}

type config struct {
	total int64
}

func NewOption(end int64) *Bar {
	current := int64(0)
	total := end
	graph := "#"
	percent := getPercent(current, total)

	return &Bar{
		percent: percent,
		current: current,
		graph:   graph,
		config: config{
			total: total,
		},
	}
}

func getPercent(current, total int64) int64 {
	return int64((float32(current) / float32(total)) * 100)
}

func (b *Bar) Add(current int) error {
	if b.config.total == 0 {
		return errors.New("the end must be greater than 0")
	}

	currentNum := int64(current)
	b.current += currentNum
	last := b.percent
	if b.current > b.config.total {
		return errors.New("current exceeds total")
	}
	b.percent = getPercent(b.current, b.config.total)
	var graphRate float64
	graphRate = 50 / float64(b.config.total)
	if b.percent != last {
		b.rate += strings.Repeat(b.graph, int(math.Ceil(graphRate)))
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", b.rate, b.percent, b.current, b.config.total)

	return nil
}

func Default(end int64) *Bar {
	return NewOption(end)
}

func main() {
	bar := Default(100)
	for i := 0; i < int(100); i++ {
		time.Sleep(10 * time.Millisecond)
		bar.Add(1)
	}
}
