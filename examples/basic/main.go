package main

import (
	"progressbar"
	"time"
)

func main() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(100 * time.Millisecond)
	}
}
