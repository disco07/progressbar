package main

import (
	"errors"
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
				bar = NewOption(tt.end + 1)
			} else {
				bar = NewOption(tt.end)
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
			bar := Default(tt.end)
			for i := 0; i <= int(tt.end); i++ {
				err := bar.Add(1)
				time.Sleep(10 * time.Millisecond)
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
