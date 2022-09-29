package main

import (
	"progressbar"
	"time"
)

func main() {
	bar := progressbar.New(100)
	bar.SetTheme(progressbar.Theme{
		GraphType:  "#",
		GraphStart: "|",
		GraphEnd:   "|",
	})
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(100 * time.Millisecond)
	}
}
