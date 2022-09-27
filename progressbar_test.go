package progressbar

import (
	"errors"
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
			expected:    errors.New("the end must be greater than 0"),
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
			description: "the end must be greater than 0",
			end:         0,
			expected:    errors.New("the end must be greater than 0"),
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
