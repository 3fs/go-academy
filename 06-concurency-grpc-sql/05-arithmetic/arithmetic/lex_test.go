package arithmetic

import "testing"
import "reflect"

func TestLex(t *testing.T) {
	cases := []struct {
		in       string
		expected []item
	}{
		{
			"1 + 23",
			[]item{
				item{typ: itemNumber, pos: 0, val: "1"},
				item{typ: itemSpace, pos: 1, val: " "},
				item{typ: itemAdding, pos: 2, val: "+"},
				item{typ: itemSpace, pos: 3, val: " "},
				item{typ: itemNumber, pos: 4, val: "23"},
				item{typ: itemEOF, pos: 6, val: ""},
			},
		},
		{
			"$1*(.5-1.2)",
			[]item{
				item{typ: itemIdentifier, pos: 0, val: "$1"},
				item{typ: itemMultiplier, pos: 2, val: "*"},
				item{typ: itemLParen, pos: 3, val: "("},
				item{typ: itemNumber, pos: 4, val: ".5"},
				item{typ: itemSubtraction, pos: 6, val: "-"},
				item{typ: itemNumber, pos: 7, val: "1.2"},
				item{typ: itemRParen, pos: 10, val: ")"},
				item{typ: itemEOF, pos: 11, val: ""},
			},
		},
	}

	for _, testCase := range cases {
		l := lex(testCase.in)

		// collext items
		items := []item{}
		for {
			item := l.nextItem()
			items = append(items, item)

			if item.typ == itemEOF || item.typ == itemError {
				break
			}
		}

		// compare it
		if !reflect.DeepEqual(items, testCase.expected) {
			t.Errorf("Issue in %s; Expected\n%+v\nto equal\n%+v", testCase.in, items, testCase.expected)
		}
	}
}
