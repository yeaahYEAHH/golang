package huffman

import (
	"archivator/lib/compression/vlc/table"
	"sort"
)

type encodingTable map[rune]string
type Generator struct{}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) NewTable(text string) table.EncodingTable {
	stat := newCharStat(text)
	leafs := buildLeafs(stat)
	tree := buildTree(leafs)
	return tree.Export()
}

type charStat map[rune]int

func newCharStat(str string) charStat {
	res := make(charStat)

	for _, c := range str {
		res[c]++
	}
	return res
}

type node struct {
	Value    rune
	Quantity int
	Zero     *node
	One      *node
}

func (n node) Export() table.EncodingTable {
	codes := make(map[rune]string)
	generateCodes(&n, "", codes)

	return codes
}

func generateCodes(node *node, path string, table map[rune]string) {
	if node.Zero == nil && node.One == nil {
		table[node.Value] = path
		return
	}

	if node.Zero != nil {
		generateCodes(node.Zero, path+"0", table)
	}
	if node.One != nil {
		generateCodes(node.One, path+"1", table)
	}
}

func buildLeafs(stat charStat) []node {
	nodes := make([]node, 0, len(stat))

	for char, quantity := range stat {
		nodes = append(nodes, node{
			Value:    char,
			Quantity: quantity,
		})
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Quantity != nodes[j].Quantity {
			return nodes[i].Quantity < nodes[j].Quantity
		}
		return nodes[i].Value < nodes[j].Value
	})

	return nodes
}

func buildTree(nodes []node) node {
	if len(nodes) == 1 {
		return nodes[0]
	}

	parent := node{
		Value:    0,
		Quantity: nodes[0].Quantity + nodes[1].Quantity,
		Zero:     &nodes[0],
		One:      &nodes[1],
	}

	nodes = nodes[2:]
	nodes = append(nodes, parent)

	return buildTree(nodes)
}
