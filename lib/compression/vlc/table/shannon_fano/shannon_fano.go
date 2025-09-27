package shannon_fano

import (
	"archivator/lib/compression/vlc/table"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Generator struct{}

type encodingTable map[rune]code

func (et encodingTable) Export() table.EncodingTable {
	res := make(table.EncodingTable)

	for k, v := range et {
		byteStr := fmt.Sprintf("%b", v.Bits)

		if lenDiff := v.Size - len(byteStr); lenDiff > 0 {
			byteStr = strings.Repeat("0", lenDiff) + byteStr
		}

		res[k] = byteStr
	}

	return res
}

type code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

type charStat map[rune]int

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	stat := newCharStat(text)
	return build(stat).Export()
}

func build(stat charStat) encodingTable {
	codes := make([]code, 0, len(stat))

	for char, quantity := range stat {
		codes = append(codes, code{
			Char:     char,
			Quantity: quantity,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}

		return codes[i].Char < codes[j].Char
	})

	assignCodes(codes)

	res := make(encodingTable)

	for _, code := range codes {
		res[code.Char] = code
	}

	return res
}

func assignCodes(codes []code) {
	if len(codes) < 2 {
		return
	}

	//	divide codes
	divider := bestDividerPosition(codes)

	// add 0 or 1
	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		codes[i].Size++

		if i >= divider {
			codes[i].Bits |= 1
		}
	}

	assignCodes(codes[:divider])
	assignCodes(codes[divider:])
}

// balance tree
func bestDividerPosition(codes []code) int {
	total := 0
	for _, code := range codes {
		total += code.Quantity
	}

	left := 0
	prevDiff := math.MaxInt
	bestPosition := 0

	for i := 0; i < len(codes)-1; i++ {
		left += codes[0].Quantity
		right := total - left
		diff := abs(right - left)

		if diff >= prevDiff {
			break
		}

		prevDiff = diff
		bestPosition = i + 1
	}

	return bestPosition
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func newCharStat(text string) charStat {
	res := make(charStat)

	for _, c := range text {
		res[c]++
	}

	return res
}
