package chunk

import (
	"reflect"
	"testing"
)

func TestSplitByChunks(t *testing.T) {
	tests := []struct {
		name      string
		bStr      string
		chunkSize int
		expected  BinaryChunks
	}{
		{
			name:     "empty string",
			bStr:     "",
			expected: BinaryChunks{},
		},
		{
			name:     "exact fit",
			bStr:     "1010101011110000",
			expected: BinaryChunks{"10101010", "11110000"},
		},
		{
			name:     "needs padding",
			bStr:     "101",
			expected: BinaryChunks{"10100000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitByChunks(tt.bStr); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name     string
		chunks   BinaryChunks
		expected string
	}{
		{
			name:     "normal binary chunks",
			chunks:   BinaryChunks{"1100", "0011", "1010"},
			expected: "110000111010",
		},
		{
			name:     "single chunk",
			chunks:   BinaryChunks{"10101010"},
			expected: "10101010",
		},
		{
			name:     "empty chunks slice",
			chunks:   BinaryChunks{},
			expected: "",
		},
		{
			name:     "chunks with empty strings",
			chunks:   BinaryChunks{"11", "", "00"},
			expected: "1100",
		},
		{
			name:     "all empty",
			chunks:   BinaryChunks{"", "", ""},
			expected: "",
		},
		{
			name:     "one empty chunk",
			chunks:   BinaryChunks{""},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.chunks.Join()
			if got != tt.expected {
				t.Errorf("BinaryChunks.Join() = %q, want %q", got, tt.expected)
			}
		})
	}
}
