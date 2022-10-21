package progressbar

import (
	"errors"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestSetTheme(t *testing.T) {
	tests := []struct {
		name      string
		parameter Theme
	}{
		{
			name: "testing graphType",
			parameter: Theme{
				GraphType:  "#",
				GraphStart: "[",
				GraphEnd:   "]",
				GraphWidth: 50,
			},
		},
		{
			name: "testing graphStart",
			parameter: Theme{
				GraphStart: "|",
				GraphType:  "█",
				GraphEnd:   "]",
				GraphWidth: 50,
			},
		},
		{
			name: "testing graphEnd",
			parameter: Theme{
				GraphStart: "[",
				GraphType:  "█",
				GraphEnd:   "|",
				GraphWidth: 50,
			},
		},
		{
			name: "testing graphWith",
			parameter: Theme{
				GraphWidth: 10,
				GraphStart: "[",
				GraphType:  "█",
				GraphEnd:   "]",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bar *Bar
			bar = New(10)
			bar.SetTheme(tt.parameter)
			if bar.theme != tt.parameter {
				t.Errorf("Bad parameter got %v want %v", bar.theme, tt.parameter)
			}
		})
	}
}

func TestNew(t *testing.T) {

	tests := []struct {
		description string
		end         int64
		expected    error
	}{
		{
			description: "work",
			end:         50,
			expected:    nil,
		},
		{
			description: "the end must be greater than 0",
			end:         0,
			expected:    errors.New("the end must be greater than zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			var bar *Bar
			if tt.description == "work" {
				bar = New(tt.end + 1)
			} else {
				bar = New(tt.end)
			}
			for i := 0; i <= int(tt.end); i++ {
				err := bar.Add(1)
				time.Sleep(100 * time.Millisecond)
				if tt.expected == nil && err != nil {
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}

				if tt.expected != nil && err.Error() != tt.expected.Error() {
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}
			}
		})
	}
}

func Test_Exceeds_Total(t *testing.T) {

	tests := []struct {
		description string
		end         int64
		expected    error
	}{
		{
			description: "current exceeds total",
			end:         10,
			expected:    errors.New("current exceeds total"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			var bar *Bar
			bar = New(tt.end)
			for i := 0; i <= int(tt.end); i++ {
				err := bar.Add(1)

				if err != nil && err.Error() != tt.expected.Error() {
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}
			}
		})
	}
}

func TestDefault(t *testing.T) {

	tests := []struct {
		description string
		end         int64
		expected    error
	}{
		{
			description: "work",
			end:         100,
			expected:    nil,
		},
		{
			description: "the end must be greater than zero",
			end:         0,
			expected:    errors.New("the end must be greater than zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			var bar *Bar
			if tt.description == "work" {
				bar = Default(tt.end + 1)
			} else {
				bar = Default(tt.end)
			}
			for i := 0; i <= int(tt.end); i++ {
				err := bar.Add(1)
				time.Sleep(100 * time.Millisecond)
				if tt.expected == nil && err != nil {
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}

				if tt.expected != nil && err.Error() != tt.expected.Error() {
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}
			}
		})
	}
}

func TestDefaultBytes(t *testing.T) {
	tests := []struct {
		desc     string
		bytes    int64
		expected error
	}{
		{
			desc:     "work",
			bytes:    100,
			expected: nil,
		},
		{
			desc:     "the end must be greater than zero",
			bytes:    0,
			expected: errors.New("the end must be greater than zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var bar *Bar
			if tt.desc == "work" {
				bar = DefaultBytes(tt.bytes + 1)
			} else {
				bar = DefaultBytes(tt.bytes)
			}
			for i := 0; i <= int(tt.bytes); i++ {
				err := bar.Add(1)
				time.Sleep(100 * time.Millisecond)
				if tt.expected == nil && err != nil {
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}

				if tt.expected != nil && err.Error() != tt.expected.Error() {
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}
			}
		})
	}
}

func TestWriter(t *testing.T) {
	req, err := http.NewRequest("GET", "https://desktop.docker.com/win/main/amd64/Docker Desktop Installer.exe", nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	f, _ := os.OpenFile("Docker Desktop Installer.exe", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := DefaultBytes(resp.ContentLength)
	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)
	if err != nil {
		t.Error(err)
	}
}
func TestReader(t *testing.T) {
	req, err := http.NewRequest("GET", "https://desktop.docker.com/win/main/amd64/Docker Desktop Installer.exe", nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	f, _ := os.OpenFile("Docker Desktop Installer.exe", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := DefaultBytes(resp.ContentLength)
	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)
	if err != nil {
		t.Error(err)
	}
}

func TestConvertTime(t *testing.T) {
	tests := []struct {
		desc     string
		time     uint
		expected string
	}{
		{
			desc:     "hour equal zero",
			time:     1000,
			expected: "16:40",
		},
		{
			desc:     "time equal zero",
			time:     0,
			expected: "00:00",
		},
		{
			desc:     "hour not equal zero",
			time:     10000,
			expected: "02:46:40",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			hour := convertTime(tt.time)

			if tt.expected != hour {
				t.Errorf("want %v got %v", tt.expected, hour)
			}
		})
	}
}

func TestUnitFormat(t *testing.T) {
	tests := []struct {
		desc     string
		it       float64
		expected string
	}{
		{
			desc:     "it >= math.Pow(1024, 4)",
			it:       10000000000000,
			expected: "9.09 TB",
		},
		{
			desc:     "it >= math.Pow(1024, 3)",
			it:       10000000000,
			expected: "9.31 GB",
		},
		{
			desc:     "it >= math.Pow(1024, 2)",
			it:       10000000,
			expected: "9.54 MB",
		},
		{
			desc:     "it >= 1024",
			it:       1024,
			expected: "1.00 KB",
		},
		{
			desc:     "it < 1024",
			it:       1000,
			expected: "1000.00 B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			unit := unitFormat(tt.it)

			if tt.expected != unit {
				t.Errorf("want %v, got %v", tt.expected, unit)
			}
		})
	}
}
