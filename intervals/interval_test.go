package intervals

import "testing"

func TestIsOverlap(t *testing.T) {
	tests := []struct {
		name     string
		a        Interval
		b        Interval
		expected bool
	}{
		{
			name:     "Overlap on the edges",
			a:        Interval{Start: 10, End: 20},
			b:        Interval{Start: 15, End: 25},
			expected: true,
		},
		{
			name:     "No overlap",
			a:        Interval{Start: 10, End: 15},
			b:        Interval{Start: 16, End: 20},
			expected: false,
		},
		{
			name:     "Complete overlap",
			a:        Interval{Start: 10, End: 30},
			b:        Interval{Start: 15, End: 25},
			expected: true,
		},
		{
			name:     "No overlap with adjacent intervals",
			a:        Interval{Start: 10, End: 15},
			b:        Interval{Start: 15, End: 20},
			expected: true,
		},
		{
			name:     "Same intervals",
			a:        Interval{Start: 10, End: 20},
			b:        Interval{Start: 10, End: 20},
			expected: true,
		},
		{
			name:     "No overlap, first ends before second starts",
			a:        Interval{Start: 10, End: 14},
			b:        Interval{Start: 15, End: 20},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOverlap(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
