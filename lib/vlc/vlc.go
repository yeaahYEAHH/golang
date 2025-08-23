package vlc

import (
	"strings"
	"unicode"

	"archivator/lib/vlc/chunk"
)

type encodingTable map[rune]string

func Encode(str string) string {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 10010101
	binStr := encodeBin(str)

	// split binary by chunks (8): bits to bytes -> 10010101 10010101 10010101
	chunks := chunk.SplitByChunks(binStr)

	// bytes to hex -> '20 30 3C'
	return chunks.ToHex().ToString()
}

func Decode(str string) string {
	hexChunks := chunk.NewHexChunks(str)

	// hexChunks -> binChunks
	binChunks := hexChunks.ToBinary()

	// binChunks -> binString
	binString := binChunks.Join()

	// build binTreeSearch

	// decode binTreeSearch

	// return build decode string
	return ""
}

func prepareText(str string) string {
	var buf strings.Builder

	for _, char := range str {
		if unicode.IsUpper(char) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(char))
		} else {
			buf.WriteRune(char)
		}
	}

	return buf.String()
}

func encodeBin(str string) string {
	var buf strings.Builder

	for _, char := range str {
		buf.WriteString(bin(char))
	}

	return buf.String()
}

func bin(char rune) string {
	table := getEncodingTable()

	res, ok := table[char]

	if !ok {
		panic("unknowa character: " + string(char))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}
