package vlc

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodingTree(t *testing.T) {

	tests := []struct {
		name string
		et   encodingTable
		want DecodeTree
	}{
		{
			name: "base tree test",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},

			want: DecodeTree{
				Zero: &DecodeTree{
					One: &DecodeTree{
						Zero: &DecodeTree{
							One: &DecodeTree{
								Value: "z",
							},
						},
					},
				},

				One: &DecodeTree{
					Zero: &DecodeTree{
						Zero: &DecodeTree{
							One: &DecodeTree{
								Value: "b",
							},
						},
					},
					One: &DecodeTree{
						Value: "a",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodeTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeTree_AddAndDecode(t *testing.T) {
	// Задаём тестовую таблицу (кодировки)
	et := encodingTable{
		'a': "0",
		'b': "10",
		'c': "11",
	}

	// Строим дерево
	tree := et.DecodeTree()

	tests := []struct {
		name string
		code string
		want string
	}{
		{"single a", "0", "a"},
		{"single b", "10", "b"},
		{"single c", "11", "c"},
		{"sequence abc", "01011", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tree.Decode(tt.code)
			if got != tt.want {
				t.Errorf("Decode(%s) = %s, want %s", tt.code, got, tt.want)
			}
		})
	}
}

func TestDecodeTree_InvalidCode(t *testing.T) {
	et := encodingTable{
		'x': "0",
		'y': "10",
	}

	tree := et.DecodeTree()

	// Вводим битовую строку, которая не соответствует коду
	got := tree.Decode("111") // нет такого пути
	want := ""                // ожидание: ничего не декодируется

	if got != want {
		t.Errorf("Decode invalid code = %s, want %s", got, want)
	}
}
