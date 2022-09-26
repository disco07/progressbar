package progressbar

import (
	"errors"
	"fmt"
	"strings"
)

type Bar struct {
	percent          float64
	current          int64
	rate             string
	graph            string
	currentGraphRate int
	config           config
}

type config struct {
	total      int64
	graphWidth int64
}

func NewOption(end int64) *Bar {
	current := int64(0)
	total := end
	graph := "█"
	percent := getPercent(current, total)

	return &Bar{
		percent: percent,
		current: current,
		graph:   graph,
		config: config{
			total:      total,
			graphWidth: 50,
		},
	}
}

func getPercent(current, total int64) float64 {
	return 100 * (float64(current) / float64(total))
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
	lastGraphRate := b.currentGraphRate
	b.currentGraphRate = int(b.percent / 100.0 * float64(b.config.graphWidth))
	if b.percent != last {
		b.rate += strings.Repeat(b.graph, b.currentGraphRate-lastGraphRate)
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", b.rate, int(b.percent), b.current, b.config.total)

	return nil
}

// Default is a basic usage of progress bar.
// In parameter, the max size of things you want to view progress.
func Default(end int64) *Bar {
	return NewOption(end)
}
