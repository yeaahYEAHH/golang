package table

import "strings"

type Generator interface {
	NewTable(text string) EncodingTable
}

type EncodingTable map[rune]string

type decodeTree struct {
	Zero  *decodeTree
	One   *decodeTree
	Value string
}

func (et EncodingTable) Decode(text string) string {
	dt := et.decodeTree()

	return dt.Decode(text)
}

func (et EncodingTable) decodeTree() decodeTree {
	res := decodeTree{}

	for char, code := range et {
		res.add(char, code)
	}

	return res
}

func (dt *decodeTree) add(value rune, code string) {
	// code:0101(0)->'z'
	currentNode := dt

	for _, char := range code {
		switch char {
		case '0':
			if currentNode.Zero == nil {
				currentNode.Zero = &decodeTree{}
			}
			currentNode = currentNode.Zero
		case '1':
			if currentNode.One == nil {
				currentNode.One = &decodeTree{}
			}
			currentNode = currentNode.One

		}
	}

	currentNode.Value = string(value)
}

func (dt *decodeTree) Decode(str string) string {
	var buf strings.Builder

	currentNode := dt

	for _, char := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = dt
		}

		switch char {
		case '0':
			currentNode = currentNode.Zero
		case '1':
			currentNode = currentNode.One

		}
	}

	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
		currentNode = dt
	}

	return buf.String()
}
