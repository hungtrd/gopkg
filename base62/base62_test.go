package xbase62

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty input",
			input:    []byte{},
			expected: "",
		},
		{
			name:     "single byte",
			input:    []byte{0x01},
			expected: "94",
		},
		{
			name:     "multiple bytes",
			input:    []byte{0x01, 0x02, 0x03},
			expected: "bhf81",
		},
		{
			name:     "longer input",
			input:    []byte{0xff, 0xee, 0xdd, 0xcc, 0xbb},
			expected: "rhL91ic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Encode(tt.input)
			if result != tt.expected {
				t.Errorf("EncodeBytes(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
	}{
		{
			name:     "empty input",
			input:    "",
			expected: []byte{},
		},
		{
			name:     "single character",
			input:    "94",
			expected: []byte{0x01},
		},
		{
			name:     "multiple characters",
			input:    "bhf81",
			expected: []byte{0x01, 0x02, 0x03},
		},
		{
			name:     "longer string",
			input:    "rhL91ic",
			expected: []byte{0xff, 0xee, 0xdd, 0xcc, 0xbb},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Decode(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("DecodeStr(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
