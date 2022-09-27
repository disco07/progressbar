package progressbar

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Bar struct {
	percent          float64
	current          int64
	rate             string
	graph            string
	currentGraphRate int
	option           Option
}

type Option struct {
	total      int64
	graphWidth int64
	startTime  time.Time
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
		option: Option{
			total:      total,
			graphWidth: 50,
			startTime:  time.Now(),
		},
	}
}

func getPercent(current, total int64) float64 {
	return 100 * (float64(current) / float64(total))
}

func (b *Bar) view() error {
	last := b.percent
	b.percent = getPercent(b.current, b.option.total)
	lastGraphRate := b.currentGraphRate
	b.currentGraphRate = int(b.percent / 100.0 * float64(b.option.graphWidth))
	if b.percent != last {
		b.rate += strings.Repeat(b.graph, b.currentGraphRate-lastGraphRate)
	}
	fmt.Printf(
		"\r[%-*s]%3d%% %5d/%d (%v)",
		b.option.graphWidth,
		b.rate,
		int(b.percent),
		b.current,
		b.option.total,
		time.Since(b.option.startTime).Round(time.Second),
	)

	return nil
}

// Add is a func who add the number passed as a parameter to the progress bar.
func (b *Bar) Add(num int) error {
	if b.option.total == 0 {
		return errors.New("the end must be greater than 0")
	}

	currentNum := int64(num)
	b.current += currentNum
	if b.current > b.option.total {
		return errors.New("current exceeds total")
	}
	b.view()
	return nil
}

// Default is a basic usage of progress bar.
// In parameter, the max size of things you want to view progress.
// It returns a pointer of Bar.
func Default(end int64) *Bar {
	return NewOption(end)
}
