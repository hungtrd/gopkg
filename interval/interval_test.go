package xinterval

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    []Interval
		expected []Interval
	}{
		{
			name:     "No intervals",
			input:    []Interval{},
			expected: []Interval{},
		},
		{
			name:     "Single interval",
			input:    []Interval{{Start: 1, End: 5}},
			expected: []Interval{{Start: 1, End: 5}},
		},
		{
			name: "Non-overlapping intervals",
			input: []Interval{
				{Start: 1, End: 3},
				{Start: 5, End: 8},
			},
			expected: []Interval{
				{Start: 1, End: 3},
				{Start: 5, End: 8},
			},
		},
		{
			name: "Overlapping intervals",
			input: []Interval{
				{Start: 1, End: 3},
				{Start: 2, End: 6},
				{Start: 8, End: 10},
			},
			expected: []Interval{
				{Start: 1, End: 6},
				{Start: 8, End: 10},
			},
		},
		{
			name: "Fully overlapping intervals",
			input: []Interval{
				{Start: 1, End: 10},
				{Start: 2, End: 5},
				{Start: 3, End: 8},
			},
			expected: []Interval{
				{Start: 1, End: 10},
			},
		},
		{
			name: "Mixed overlapping and non-overlapping intervals",
			input: []Interval{
				{Start: 1, End: 3},
				{Start: 4, End: 6},
				{Start: 5, End: 7},
				{Start: 8, End: 10},
			},
			expected: []Interval{
				{Start: 1, End: 3},
				{Start: 4, End: 7},
				{Start: 8, End: 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}

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
