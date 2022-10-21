package progressbar

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

type Bar struct {
	state  state
	option option
	theme  Theme
}

type state struct {
	percent          float64
	current          int64
	currentGraphRate int
}

type Theme struct {
	rate       string
	GraphType  string
	GraphStart string
	GraphEnd   string
	GraphWidth int64
}

type option struct {
	total     int64
	startTime time.Time
	bytes     bool
	sync.Mutex
}

func (b *Bar) SetTheme(t Theme) {
	if t.GraphType != "" {
		b.theme.GraphType = t.GraphType
	}
	if t.GraphWidth != 0 {
		b.theme.GraphWidth = t.GraphWidth
	}
	if t.GraphStart != "" {
		b.theme.GraphStart = t.GraphStart
	}
	if t.GraphEnd != "" {
		b.theme.GraphEnd = t.GraphEnd
	}
}

func New(end int64) *Bar {
	return &Bar{
		state: state{
			percent: getPercent(int64(0), end),
			current: int64(0),
		},
		theme: Theme{
			GraphType:  "█",
			GraphStart: "|",
			GraphEnd:   "|",
			GraphWidth: 60,
		},
		option: option{
			total:     end,
			startTime: time.Now(),
			bytes:     false,
		},
	}
}

func getPercent(current, total int64) float64 {
	return 100 * (float64(current) / float64(total))
}

func (b *Bar) view() error {
	// iteration per second
	var itPerS float64
	// convert it to seconds in some format
	var itUnits string
	last := b.state.percent
	b.state.percent = getPercent(b.state.current, b.option.total)
	lastGraphRate := b.state.currentGraphRate
	b.state.currentGraphRate = int(b.state.percent / 100.0 * float64(b.theme.GraphWidth))
	if b.state.percent != last {
		b.theme.rate += strings.Repeat(b.theme.GraphType, b.state.currentGraphRate-lastGraphRate)
	}

	timeElapsed := uint(time.Since(b.option.startTime).Seconds())
	timeLeft := uint(time.Since(b.option.startTime).Seconds() / float64(b.state.current) * (float64(b.option.total) - float64(b.state.current)))

	if timeElapsed >= 1 {
		itPerS = float64(uint(b.state.current) / timeElapsed)
	}

	if b.option.bytes {
		itUnits = unitFormat(itPerS)
	} else {
		itUnits = fmt.Sprintf("%v it", itPerS)
	}

	fmt.Printf(
		"\r %3d%% %s%-*s%s [%v-%v, %v, %d/%d]  ",
		int(b.state.percent),
		b.theme.GraphStart,
		b.theme.GraphWidth,
		b.theme.rate,
		b.theme.GraphEnd,
		convertTime(timeElapsed),
		convertTime(timeLeft),
		itUnits,
		b.state.current,
		b.option.total,
	)

	return nil
}

// Add is a func who add the number passed as a parameter to the progress bar.
func (b *Bar) Add(num int) error {
	b.option.Lock()
	defer b.option.Unlock()
	if b.option.total == 0 {
		return errors.New("the end must be greater than zero")
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
	return New(end)
}

func DefaultBytes(end int64) *Bar {
	return &Bar{
		state: state{
			percent: getPercent(int64(0), end),
			current: int64(0),
		},
		theme: Theme{
			GraphType:  "█",
			GraphStart: "|",
			GraphEnd:   "|",
			GraphWidth: 60,
		},
		option: option{
			total:     end,
			startTime: time.Now(),
			bytes:     true,
		},
	}
}

func convertTime(second uint) string {
	var seconds = second % 60
	var minutes = (second / 60) % 60
	var hours = (second / 60) / 60
	if hours == 0 {
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func unitFormat(it float64) string {
	var kiloBytes = 1024.0
	if it >= math.Pow(kiloBytes, 4) {
		return fmt.Sprintf("%0.2f TB", it/math.Pow(kiloBytes, 4))
	} else if it >= math.Pow(kiloBytes, 3) {
		return fmt.Sprintf("%0.2f GB", it/math.Pow(kiloBytes, 3))
	} else if it >= math.Pow(kiloBytes, 2) {
		return fmt.Sprintf("%0.2f MB", it/math.Pow(kiloBytes, 2))
	} else if it >= kiloBytes {
		return fmt.Sprintf("%0.2f KB", it/kiloBytes)
	}

	return fmt.Sprintf("%0.2f B", it)
}
