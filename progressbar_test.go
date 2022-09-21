package progressbar

import (
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
			description: "normal progress bar",
			end:         0,
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			bar := NewOption(tt.end)
			for i := 0; i < int(tt.end); i++ {
				if err := bar.PlayBar(i); err != nil {
					time.Sleep(10 * time.Millisecond)
					t.Errorf("got %v want %v", err.Error(), tt.expected.Error())
				}
			}
		})
	}
}
