### Basic usage
```golang
package main

import (
	"github.com/disco07/progressbar"
	"time"
)

func main() {
	bar := progressbar.New(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(100 * time.Millisecond)
	}
}
```
![Basic bar](progressbar.gif)
