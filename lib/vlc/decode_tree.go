package vlc

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
				currentNode = &DecodeTree{}
			}
			currentNode = currentNode.Zero
		case '1':
			if currentNode.One == nil {
				currentNode = &DecodeTree{}
			}
			currentNode = currentNode.One

		}
	}

	currentNode.Value = string(value)
}
