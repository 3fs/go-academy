package arithmetic

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type (
	// identifier for lex items
	itemType int

	// pos represents position inside a string
	pos int

	// stateFn represents the state of the scanner as a function that returns
	// the next state
	stateFn func(l *lexer) stateFn

	// item represents a single token
	item struct {
		typ itemType
		pos pos
		val string
	}

	// lexer holds the state of the scanner
	lexer struct {
		input string
		state stateFn
		width pos
		start pos
		pos   pos
		items chan item
	}
)

const (
	// EOF marker
	eof = -1

	itemError       itemType = iota // error occured
	itemEOF                         // EOF helper
	itemSpace                       // whitespace
	itemLParen                      // left parenthesis
	itemRParen                      // right parenthesis
	itemNumber                      // numeric value
	itemMultiplier                  // multiplication character (*)
	itemDivision                    // division character (/)
	itemAdding                      // adding character (+)
	itemSubtraction                 // subtraction character (-)
	itemIdentifier                  // placeholder for a previous result ($1)
	itemUnknown                     // placeholder for a previous result ($1)
)

func (i item) Exec(st state) state {
	switch i.typ {
	case itemNumber:
		// try to convert the string to float
		f, err := strconv.ParseFloat(i.val, 64)
		if err != nil {
			st.err = newError(st.input, fmt.Sprintf("Failed to convert %s to float; %v", i.val, err), i)
			return st
		}

		st.currentValue = f
	case itemIdentifier:
		num, err := strconv.ParseInt(i.val[1:], 10, 32)
		if err != nil {
			st.err = newError(st.input, fmt.Sprintf("Failed to extract number from identifier \"%s\"; %v", i.val, err), i)
			return st
		}

		val, err := st.storage.Get(int(num))
		if err != nil {
			st.err = newError(st.input, "Failed to read value from storage", i)
			return st
		}

		st.currentValue = val

	default:
		st.err = newError(st.input, fmt.Sprintf("Identifier does not support Exec"), i)
	}

	return st
}

func lex(in string) *lexer {
	l := &lexer{
		input: in,
		items: make(chan item),
	}

	go l.run()

	return l
}

func (l *lexer) run() {
	for l.state = lexSection; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	// fmt.Println("next", l.pos, len(l.input), l.input)
	if int(l.pos) >= len(l.input) {
		// fmt.Println("next", "eof")
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = pos(w)
	l.pos += l.width
	// fmt.Println("next", "width", l.width, r)
	return r
}

func (l *lexer) nextItem() item {
	return <-l.items
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	// fmt.Println("emit", t, l.start, l.input[l.start:l.pos])
	l.items <- item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

func lexSection(l *lexer) stateFn {
	switch r := l.next(); {
	case r == eof:
		l.emit(itemEOF)
		return nil
	case isSpace(r):
		return lexSpace
	case r == '(':
		l.emit(itemLParen)
	case r == ')':
		l.emit(itemRParen)
	case r == '*':
		l.emit(itemMultiplier)
	case r == '/':
		l.emit(itemDivision)
	case r == '-':
		l.emit(itemSubtraction)
	case r == '+':
		l.emit(itemAdding)
	case isNumeric(r):
		return lexNumber
	case r == '$':
		return lexIdentifier
	default:
		l.emit(itemError)
	}

	return lexSection
}

// lexSpace scans a run of space characters. One space has already been seen.
func lexSpace(l *lexer) stateFn {
	for isSpace(l.peek()) {
		l.next()
	}
	l.emit(itemSpace)
	return lexSection
}

// lexNumber scans a number
func lexNumber(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isNumeric(r):
			// absorb.
		default:
			l.backup()
			l.emit(itemNumber)
			break Loop
		}
	}

	return lexSection
}

// lexIdentifier scans an identifier
func lexIdentifier(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case unicode.IsDigit(r):
			// absorb
		default:
			l.backup()
			l.emit(itemIdentifier)
			break Loop
		}
	}

	return lexSection
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isNumeric reports whether r is an digit or a dot
func isNumeric(r rune) bool {
	return r == '.' || unicode.IsDigit(r)
}
