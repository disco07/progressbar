package progressbar

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestNewOption(t *testing.T) {

	tests := []struct {
		description string
		end         int64
		expected    error
	}{
		{
			description: "normal progress bar",
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
			bar := newOption(tt.end)
			for i := 0; i < int(tt.end); i++ {
				if err := bar.PlayBar(i); err != nil {
					time.Sleep(10 * time.Millisecond)
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
			description: "normal progress bar",
			end:         100,
			expected:    nil,
		},
		{
			description: "the end must be greater than 0",
			end:         0,
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			bar := Default(tt.end)
			for i := 0; i < int(tt.end); i++ {
				err := bar.PlayBar(i)
				fmt.Println(err)
				if err != nil {
					time.Sleep(10 * time.Millisecond)
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}
			}
		})
	}
}
