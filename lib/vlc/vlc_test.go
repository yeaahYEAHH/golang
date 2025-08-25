package vlc

import (
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

func TestEncode(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
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

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "20 30 3C 18 77 4A E4 4D 28",
			want: "My name is Ted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
