package vlc

import (
	"reflect"
	"testing"
)

func TestPrepareText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "lowercase only",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "uppercase only",
			input:    "HELLO",
			expected: "!h!e!l!l!o",
		},
		{
			name:     "mixed case",
			input:    "HeLLo WoRLd",
			expected: "!he!l!lo !wo!r!ld",
		},
		{
			name:     "single uppercase",
			input:    "A",
			expected: "!a",
		},
		{
			name:     "single lowercase",
			input:    "a",
			expected: "a",
		},
		{
			name:     "numbers and symbols",
			input:    "Test123!@#",
			expected: "!test123!@#",
		},
		{
			name:     "spaces and tabs",
			input:    "Hi There",
			expected: "!hi !there",
		},
		{
			name:     "non-ASCII letters (Cyrillic)",
			input:    "Привет МИР",
			expected: "!привет !м!и!р",
		},
		{
			name:     "emoji (not letters)",
			input:    "Hello 🌍",
			expected: "!hello 🌍",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prepareText(tt.input)
			if got != tt.expected {
				t.Errorf("prepareText(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGetEncodingTable(t *testing.T) {
	table := getEncodingTable()

	tests := map[rune]string{
		'e': "101",
		't': "1001",
		'a': "011",
		'o': "10001",
		'n': "10000",
		's': "0101",
		'!': "001000",
		' ': "11",
		'l': "001001",
		'z': "000000000000",
	}

	for char, expected := range tests {
		if got, ok := table[char]; !ok {
			t.Errorf("getEncodingTable() missing encoding for '%c'", char)
		} else if got != expected {
			t.Errorf("getEncodingTable()['%c'] = %q, want %q", char, got, expected)
		}
	}
}

func TestBin(t *testing.T) {
	tests := []struct {
		char     rune
		expected string
	}{
		{'e', "101"},
		{'t', "1001"},
		{'a', "011"},
		{' ', "11"},
		{'!', "001000"},
		{'z', "000000000000"},
	}

	for _, tt := range tests {
		t.Run(string(tt.char), func(t *testing.T) {
			got := bin(tt.char)
			if got != tt.expected {
				t.Errorf("bin('%c') = %q, want %q", tt.char, got, tt.expected)
			}
		})
	}

	t.Run("unknown character", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("ожидалась паника при неизвестном символе, но её не было")
			}
		}()
		bin('Л')
	})
}

func TestEncodeBin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single char",
			input:    "e",
			expected: "101",
		},
		{
			name:     "hello lowercase",
			input:    "hello",
			expected: "001110100100100100110001",
		},
		{
			name:     "space and e",
			input:    " e",
			expected: "11101",
		},
		{
			name:     "prepared uppercase: !h",
			input:    "!h",
			expected: "0010000011",
		},
		{
			name:     "complex: !m!a!n",
			input:    "!m!a!n",
			expected: "00100000001100100001100100010000",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeBin(tt.input)
			if got != tt.expected {
				t.Errorf("encodeBin(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestSplitByChunks(t *testing.T) {
	tests := []struct {
		name      string
		bStr      string
		chunkSize int
		expected  BinaryChunks
	}{
		{
			name:      "empty string",
			bStr:      "",
			chunkSize: 8,
			expected:  BinaryChunks{},
		},
		{
			name:      "exact fit",
			bStr:      "1010101011110000",
			chunkSize: 8,
			expected:  BinaryChunks{"10101010", "11110000"},
		},
		{
			name:      "needs padding",
			bStr:      "101",
			chunkSize: 8,
			expected:  BinaryChunks{"10100000"},
		},
		{
			name:      "multiple chunks with padding",
			bStr:      "11001100110",
			chunkSize: 4,
			expected:  BinaryChunks{"1100", "1100", "1100"},
		},
		{
			name:      "chunk size 1",
			bStr:      "101",
			chunkSize: 1,
			expected:  BinaryChunks{"1", "0", "1"},
		},
		{
			name:      "chunk size larger than string",
			bStr:      "10",
			chunkSize: 5,
			expected:  BinaryChunks{"10000"},
		},
		{
			name:      "chunk size 0 (edge case)",
			bStr:      "1010",
			chunkSize: 0,
			expected:  BinaryChunks{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.bStr, tt.chunkSize); !reflect.DeepEqual(got, tt.expected) {
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
			sep:      ",",
			expected: "",
		},
		{
			name:     "single chunk, no separator",
			hcs:      HexChunks{"AA"},
			sep:      ",",
			expected: "AA",
		},
		{
			name:     "two chunks with comma",
			hcs:      HexChunks{"AA", "FF"},
			sep:      ",",
			expected: "AA,FF",
		},
		{
			name:     "multiple chunks with comma",
			hcs:      HexChunks{"10", "2A", "B3", "01"},
			sep:      ",",
			expected: "10,2A,B3,01",
		},
		{
			name:     "two chunks with space",
			hcs:      HexChunks{"AA", "FF"},
			sep:      " ",
			expected: "AA FF",
		},
		{
			name:     "multiple chunks with custom separator",
			hcs:      HexChunks{"A", "B", "C"},
			sep:      " -> ",
			expected: "A -> B -> C",
		},
		{
			name:     "empty separator",
			hcs:      HexChunks{"01", "AB", "CD"},
			sep:      "",
			expected: "01ABCD",
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
			got := tt.hcs.toString(tt.sep)
			if got != tt.expected {
				t.Errorf("toString(%v, %q) = %q, want %q", tt.hcs, tt.sep, got, tt.expected)
			}
		})
	}
}

func TestEncode(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "20303C18774AE44D28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
