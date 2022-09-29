[![test](https://github.com/disco07/progressbar/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/disco07/progressbar/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/disco07/progressbar.svg)](https://pkg.go.dev/github.com/disco07/progressbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/disco07/progressbar)](https://goreportcard.com/report/github.com/disco07/progressbar)
[![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen.svg)](https://gocover.io/github.com/disco07/progressbar)

# progressbar
A simple progress bar for golang projects. I created that because there was a long processus in some projet and I didn't know what was the progression.


## Installation
```golang
go get -u github.com/disco07/progressbar
```

## Usage
### Basic usage
```golang
bar := progressbar.Default(100)
for i := 0; i < 100; i++ {
    bar.Add(1)
    time.Sleep(100 * time.Millisecond)
}
```
![Basic bar](examples/basic/progressbar.gif)

## Contributing 🤝
Contributions, issues, and feature requests are welcome!

Feel free to check the issues page.

## 📝 License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
