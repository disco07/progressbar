package progressbar

import (
	"errors"
	"fmt"
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

func (b *Bar) PlayBar(current int) error {
	if b.total == 0 {
		return errors.New("the end must be greater than 0")
	}

	currentNum := int64(current)
	b.current = currentNum
	last := b.percent
	b.percent = getPercent(currentNum, b.total)
	if b.percent != last && b.percent%2 == 0 {
		b.rate += b.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", b.rate, b.percent, b.current, b.total)

	return nil
}

func Default(end int64) *Bar {
	return NewOption(end)
}
