package progressbar

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Bar struct {
	rate   string
	state  State
	option Option
	theme  Theme
}

type State struct {
	percent          float64
	current          int64
	currentGraphRate int
}

type Theme struct {
	Graph      string
	GraphWidth int64
}

type Option struct {
	total     int64
	startTime time.Time
}

func (b *Bar) SetTheme(t Theme) {
	if t.Graph != "" {
		b.theme.Graph = t.Graph
	}
	if t.GraphWidth != 0 {
		b.theme.GraphWidth = t.GraphWidth
	}
}

func NewOption(end int64) *Bar {
	current := int64(0)
	total := end
	graph := "█"
	percent := getPercent(current, total)

	return &Bar{
		state: State{
			percent: percent,
			current: current,
		},
		theme: Theme{
			Graph:      graph,
			GraphWidth: 50,
		},
		option: Option{
			total:     total,
			startTime: time.Now(),
		},
	}
}

func getPercent(current, total int64) float64 {
	return 100 * (float64(current) / float64(total))
}

func (b *Bar) view() error {
	last := b.state.percent
	b.state.percent = getPercent(b.state.current, b.option.total)
	lastGraphRate := b.state.currentGraphRate
	b.state.currentGraphRate = int(b.state.percent / 100.0 * float64(b.theme.GraphWidth))
	if b.state.percent != last {
		b.rate += strings.Repeat(b.theme.Graph, b.state.currentGraphRate-lastGraphRate)
	}
	secondsLeft := time.Since(b.option.startTime).Seconds() / float64(b.state.current) * (float64(b.option.total) - float64(b.state.current))
	fmt.Printf(
		"\r[%-*s]%3d%% %4d/%d (%v-%v)",
		b.theme.GraphWidth,
		b.rate,
		int(b.state.percent),
		b.state.current,
		b.option.total,
		time.Since(b.option.startTime).Round(time.Second),
		time.Duration(secondsLeft)*time.Second,
	)

	return nil
}

// Add is a func who add the number passed as a parameter to the progress bar.
func (b *Bar) Add(num int) error {
	if b.option.total == 0 {
		return errors.New("the end must be greater than 0")
	}

	currentNum := int64(num)
	b.state.current += currentNum
	if b.state.current > b.option.total {
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
