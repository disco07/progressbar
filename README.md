[![test](https://github.com/disco07/progressbar/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/disco07/progressbar/actions/workflows/test.yml)
![Coverage](https://img.shields.io/badge/Coverage-95.7%25-brightgreen)

# progressbar

## Installation
```bash
go get -u github.com/disco07/progressbar
```

## Usage
### Basic usage
```bash
bar := progressbar.Default(100)
for i := 0; i < 100; i++ {
    bar.Add(1)
    time.Sleep(100 * time.Millisecond)
}
```
returns

![Alt Text](https://github.com/disco07/progressbar/Example/basic/progressbar.gif)

## Contributing 🤝
Contributions, issues, and feature requests are welcome!

Feel free to check the issues page.

## 📝 License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)