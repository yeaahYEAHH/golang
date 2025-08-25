package vlc

import "strings"

type DecodeTree struct {
	Zero  *DecodeTree
	One   *DecodeTree
	Value string
}

func (et encodingTable) DecodeTree() DecodeTree {
	res := DecodeTree{}

	for char, code := range et {
		res.Add(char, code)
	}

	return res
}

func (dt *DecodeTree) Add(value rune, code string) {
	// code:0101(0)->'z'
	currentNode := dt

	for _, char := range code {
		switch char {
		case '0':
			if currentNode.Zero == nil {
				currentNode.Zero = &DecodeTree{}
			}
			currentNode = currentNode.Zero
		case '1':
			if currentNode.One == nil {
				currentNode.One = &DecodeTree{}
			}
			currentNode = currentNode.One

		}
	}

	currentNode.Value = string(value)
}

func (dt *DecodeTree) Decode(str string) string {
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
