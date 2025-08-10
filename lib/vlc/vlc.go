package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable map[rune]string

type BinaryChunk string
type BinaryChunks []BinaryChunk

type HexChunk string
type HexChunks []HexChunk

var chunkSize = 8

func Encode(str string) string {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 10010101
	binStr := encodeBin(str)

	// split binary by chunks (8): bits to bytes -> 10010101 10010101 10010101
	chunks := splitByChunks(binStr, chunkSize)

	// bytes to hex -> '20 30 3C'
	return chunks.ToHex().toString("")
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

func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	if chunkSize <= 0 {
		return BinaryChunks{}
	}

	strLen := utf8.RuneCountInString(bStr)

	chunkCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunkCount++
	}

	res := make(BinaryChunks, 0, chunkCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		res = append(res, chunk.ToHex())
	}

	return res
}

func (bc BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("can`t parse binaty chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func (hcs HexChunks) toString(sep string) string {
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, char := range hcs[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(char))
	}

	return buf.String()
}
