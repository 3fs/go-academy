package arithmetic

import (
	"fmt"
)

type state struct {
	input        string
	err          error
	storage      storage
	currentValue float64
}

type storage interface {
	Get(int) (float64, error)
}

type calculator interface {
	Exec(state) state
}

type operation struct {
	op  item
	val calculator
}

type calculation struct {
	items []calculator
}

type parseError struct {
	input string
	msg   string
	item  item
}

func Calculate(in string, st storage) (float64, error) {
	c, err := parse(in)
	if err != nil {
		return 0, err
	}

	result := c.Exec(state{input: in, storage: st})
	return result.currentValue, result.err
}

func newError(in, msg string, i item) *parseError {
	return &parseError{
		input: in,
		msg:   msg,
		item:  i,
	}
}

func (e parseError) Error() string {
	return fmt.Sprintf("%s^%s: Failed to parse calculation; err = %v", e.input[:e.item.pos], e.input[e.item.pos:], e.msg)
}

func parse(in string) (*calculation, error) {
	l := lex(in)
	return parseInput(l, 0)
}

func parseInput(l *lexer, depth int) (*calculation, error) {
	c := &calculation{items: []calculator{}}
	var op *operation

	itemNum := 0
Loop:
	for {
		itemNum++

		switch i := l.nextItem(); {
		case i.typ == itemSpace:
			// skip spaces
			continue

		case i.typ == itemEOF:
			// handle EOF

			// check if any parentheses were left open
			if depth > 0 {
				return nil, newError(l.input, "Premature end reached, parentheses not closed", i)
			}

			// check if operation was not closed correctly
			if op != nil {
				return nil, newError(l.input, "Premature end reached, operation not closed", i)
			}

			break Loop

		case i.typ == itemError:
			// error found
			return nil, newError(l.input, "Failed to parse", i)

		case isOperator(i.typ):
			// operator reached

			// check if operation is already initialized
			if op != nil {
				return nil, newError(l.input, "Unexpected operator reached", i)
			}

			// initialize a new operation
			op = &operation{op: i}

			// skip to next item
			continue

		case i.typ == itemLParen:
			// left parentheses reached, start a new calculation
			c, err := parseInput(l, depth+1)
			if err != nil {
				return nil, err
			}

			// are we dealing with a first item
			if itemNum == 1 {
				c.items = append(c.items, c)
				continue
			}

			// update the operation
			op.val = c

		case i.typ == itemRParen:
			// right parentheses, close the current calculation
			return c, nil

		case i.typ == itemNumber || i.typ == itemIdentifier:
			// reached number or an identifier

			// handle first item, don't create an operation
			if itemNum == 1 {
				c.items = append(c.items, i)
				continue
			}

			if op == nil {
				return nil, newError(l.input, "Unexpected value found", i)
			}

			// update the operation
			op.val = i
		default:
		}

		// add the operation to the list
		c.items = append(c.items, op)

		// clear the operation
		op = nil
	}

	return c, nil
}

// Exec handles the operation and does the actual math depending on the operator
// type.
func (o operation) Exec(st state) state {
	// calculate the value to operate with
	st2 := o.val.Exec(state{input: st.input, storage: st.storage})
	if st2.err != nil {
		return st2
	}

	// do the math
	switch o.op.typ {
	case itemAdding:
		st.currentValue += st2.currentValue
	case itemSubtraction:
		st.currentValue -= st2.currentValue
	case itemMultiplier:
		st.currentValue *= st2.currentValue
	case itemDivision:
		// check for division by zero
		if st2.currentValue == 0 {
			st.err = fmt.Errorf("Division by zero")
		}
		st.currentValue /= st2.currentValue
	}

	return st
}

// Exec iterates over all the items in the calculation and executes all
// operations on them.
func (c calculation) Exec(st state) state {
	st = state{input: st.input, storage: st.storage}

	for _, i := range c.items {
		st = i.Exec(st)

		// stop if error reached
		if st.err != nil {
			return st
		}
	}

	return st
}

// isOperator checks if the type belongs to one of the operators.
func isOperator(it itemType) bool {
	return it == itemAdding || it == itemSubtraction || it == itemDivision || it == itemMultiplier
}
