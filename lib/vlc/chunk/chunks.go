package chunk

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunk string
type BinaryChunks []BinaryChunk

type HexChunk string
type HexChunks []HexChunk

const chunkSize = 8
const separator = " "

func SplitByChunks(bStr string) BinaryChunks {
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

func (bcs BinaryChunks) Join() string {
	var res strings.Builder

	for _, bc := range bcs {
		res.WriteString(string(bc))
	}

	return res.String()
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

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, separator)

	res := make(HexChunks, 0, len(parts))

	for _, chunk := range parts {
		res = append(res, HexChunk(chunk))
	}

	return res
}

func (hcs HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hcs))

	for _, chunk := range hcs {
		res = append(res, chunk.ToBinary())
	}

	return res
}

func (hc HexChunk) ToBinary() BinaryChunk {
	num, err := strconv.ParseUint(string(hc), 16, chunkSize)

	if err != nil {
		panic("can`t parse binaty chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%08b", num))

	return BinaryChunk(res)
}

func (hcs HexChunks) ToString() string {
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hcs[0]))

	for _, char := range hcs[1:] {
		buf.WriteString(separator)
		buf.WriteString(string(char))
	}

	return buf.String()
}
