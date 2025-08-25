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

func TestBinaryChunk_ToHex(t *testing.T) {
	tests := []struct {
		name     string
		bc       BinaryChunk
		expected HexChunk
	}{
		{
			name:     "zero byte",
			bc:       "00000000",
			expected: "00",
		},
		{
			name:     "max byte",
			bc:       "11111111",
			expected: "FF",
		},
		{
			name:     "AA pattern",
			bc:       "10101010",
			expected: "AA",
		},
		{
			name:     "55 pattern",
			bc:       "01010101",
			expected: "55",
		},
		{
			name:     "single bit",
			bc:       "00000001",
			expected: "01",
		},
		{
			name:     "7F",
			bc:       "01111111",
			expected: "7F",
		},
		{
			name:     "80",
			bc:       "10000000",
			expected: "80",
		},
		{
			name:     "short chunk (padded internally by logic)",
			bc:       "101",
			expected: "05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bc.ToHex()
			if got != tt.expected {
				t.Errorf("ToHex(%q) = %q, want %q", tt.bc, got, tt.expected)
			}
		})
	}

	t.Run("invalid binary", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ожидалась паника при некорректной двоичной строке")
			}
		}()
		_ = BinaryChunk("102010").ToHex()
	})
}

func TestBinaryChunks_ToHex(t *testing.T) {
	input := BinaryChunks{
		"10101010",
		"11110000",
		"00001111",
		"00000001",
	}
	expected := HexChunks{"AA", "F0", "0F", "01"}

	got := input.ToHex()

	if len(got) != len(expected) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], expected[i])
		}
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

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected HexChunks
	}{
		{
			name:     "normal comma-separated string",
			str:      "a1 b2 c3",
			expected: HexChunks{"a1", "b2", "c3"},
		},
		{
			name:     "single value",
			str:      "abc123",
			expected: HexChunks{"abc123"},
		},
		{
			name:     "empty string",
			str:      "",
			expected: HexChunks{""},
		},
		{
			name:     "string with empty parts",
			str:      "a  b",
			expected: HexChunks{"a", "", "b"},
		},
		{
			name:     "only separator",
			str:      "  ",
			expected: HexChunks{"", "", ""},
		},
		{
			name:     "spaces and symbols",
			str:      "ff, aa, 10",
			expected: HexChunks{"ff,", "aa,", "10"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NewHexChunks(%q) = %v, want %v", tt.str, got, tt.expected)
			}
		})
	}
}

func TestHexChunks_toString(t *testing.T) {
	tests := []struct {
		name     string
		hcs      HexChunks
		sep      string
		expected string
	}{
		{
			name:     "empty list",
			hcs:      HexChunks{},
			expected: "",
		},
		{
			name:     "single chunk",
			hcs:      HexChunks{"AA"},
			expected: "AA",
		},
		{
			name:     "two chunks",
			hcs:      HexChunks{"AA", "FF"},
			expected: "AA FF",
		},
		{
			name:     "multiple chunks",
			hcs:      HexChunks{"10", "2A", "B3", "01"},
			expected: "10 2A B3 01",
		},
		{
			name:     "nil slice",
			hcs:      nil,
			sep:      ",",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hcs.ToString()
			if got != tt.expected {
				t.Errorf("toString(%v, %q) = %q, want %q", tt.hcs, tt.sep, got, tt.expected)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name        string
		hexChunks   HexChunks
		wantBinary  BinaryChunks
		expectPanic bool
	}{
		{
			name:        "valid chunks: ff, 00, a5",
			hexChunks:   HexChunks{"ff", "00", "a5"},
			wantBinary:  BinaryChunks{"11111111", "00000000", "10100101"},
			expectPanic: false,
		},
		{
			name:        "single chunk",
			hexChunks:   HexChunks{"7"},
			wantBinary:  BinaryChunks{"00000111"},
			expectPanic: false,
		},
		{
			name:        "contains invalid hex",
			hexChunks:   HexChunks{"aa", "gg", "bb"},
			expectPanic: true,
		},
		{
			name:        "empty chunks slice",
			hexChunks:   HexChunks{},
			wantBinary:  BinaryChunks{},
			expectPanic: false,
		},
		{
			name:        "chunk with empty string",
			hexChunks:   HexChunks{"ff", "", "aa"},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			got := tt.hexChunks.ToBinary()

			if tt.expectPanic {
				t.Fatal("expected panic, but execution continued")
			}

			if !reflect.DeepEqual(got, tt.wantBinary) {
				t.Errorf("HexChunks.ToBinary() = %v, want %v", got, tt.wantBinary)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name       string
		hex        HexChunk
		wantBinary BinaryChunk
		wantPanic  bool
	}{
		{
			name:       "valid hex ff -> 11111111",
			hex:        "ff",
			wantBinary: "11111111",
			wantPanic:  false,
		},
		{
			name:       "valid hex 00 -> 00000000",
			hex:        "00",
			wantBinary: "00000000",
			wantPanic:  false,
		},
		{
			name:       "valid hex a5 -> 10100101",
			hex:        "a5",
			wantBinary: "10100101",
			wantPanic:  false,
		},
		{
			name:       "single digit 7 -> 00000111",
			hex:        "7",
			wantBinary: "00000111",
			wantPanic:  false,
		},
		{
			name:      "empty string",
			hex:       "",
			wantPanic: true,
		},
		{
			name:      "invalid hex: contains 'g'",
			hex:       "fg",
			wantPanic: true,
		},
		{
			name:      "invalid hex: special chars",
			hex:       "a!",
			wantPanic: true,
		},
		{
			name:       "uppercase hex should work",
			hex:        "AA",
			wantBinary: "10101010",
			wantPanic:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			got := tt.hex.ToBinary()

			if tt.wantPanic {
				t.Error("expected panic, but did not panic")
			}

			if !tt.wantPanic && got != tt.wantBinary {
				t.Errorf("HexChunk(%q).ToBinary() = %q, want %q", tt.hex, got, tt.wantBinary)
			}
		})
	}
}
