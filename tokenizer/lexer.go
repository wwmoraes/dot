package tokenizer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	// RuneEOF fake rune for end-of-file cases
	RuneEOF rune = -1
)

// Lexer is implemented by string lexical analyzers
type Lexer interface {
	hasPrefix(string) bool
	hasMoved() bool
	emit(tokenType TokenType)
	ignore()
	backup()
	restart()
	advance(int)
	next() rune
	peek() rune
	accept(string) bool
	acceptExcept(string) bool
	consume(string)
	consumeExcept(string)
	errorf(string, ...interface{}) stateFunctor
	current() string
	skip(string)
}

type lexer struct {
	input    string
	start    int
	width    int
	position int
	tokens   chan Token
}

// hasPrefix checks if the input starts with search from the current position
func (thisLexer *lexer) hasPrefix(search string) bool {
	return strings.HasPrefix(thisLexer.input[thisLexer.position:], search)
}

// hasMoved returns if the current position is ahead of the start
func (thisLexer *lexer) hasMoved() bool {
	return thisLexer.position > thisLexer.start
}

// emit sends a Token to the channel and moves the start point forward
func (thisLexer *lexer) emit(tokenType TokenType) {
	thisLexer.tokens <- Token{
		Type:  tokenType,
		Value: thisLexer.current(),
	}

	// "ignore" i.e. move the start to current position and reset width
	thisLexer.ignore()
}

// ignore moves the start point to the current position without using the data
func (thisLexer *lexer) ignore() {
	thisLexer.start = thisLexer.position
	thisLexer.width = 0
}

// backup restores position prior to the last advance/next call
func (thisLexer *lexer) backup() {
	thisLexer.position -= thisLexer.width
	thisLexer.width = 0
}

// restart restores the position prior to the last emit
func (thisLexer *lexer) restart() {
	thisLexer.position = thisLexer.start
	thisLexer.width = 0
}

// advance moves the position forward by an arbitrary amount
func (thisLexer *lexer) advance(length int) {
	thisLexer.position += length
	thisLexer.width = length
}

// next moves the position forward by one rune and return it, or EOF
func (thisLexer *lexer) next() (char rune) {
	if thisLexer.position >= len(thisLexer.input) {
		thisLexer.width = 0
		return RuneEOF
	}

	char, thisLexer.width = utf8.DecodeRuneInString(thisLexer.input[thisLexer.position:])
	thisLexer.position += thisLexer.width

	return char
}

// peek looks ahead and returns the next rune while keeping the current position
func (thisLexer *lexer) peek() rune {
	char := thisLexer.next()
	thisLexer.backup()
	return char
}

// accept moves position to the next rune if it is present on the given string
func (thisLexer *lexer) accept(chars string) bool {
	if strings.ContainsRune(chars, thisLexer.next()) {
		return true
	}
	thisLexer.backup()
	return false
}

// acceptExcept moves position to the next rune if it is not present on the
// given string
func (thisLexer *lexer) acceptExcept(chars string) bool {
	if !strings.ContainsRune(chars, thisLexer.next()) {
		return true
	}
	thisLexer.backup()
	return false
}

// consume moves the position forward for all runes that match the given string
func (thisLexer *lexer) consume(chars string) {
	for strings.ContainsRune(chars, thisLexer.next()) {
	}
	thisLexer.backup()
}

// consumeExcept moves the position forward for all runes that do not match the
// given string
func (thisLexer *lexer) consumeExcept(chars string) {
	for !strings.ContainsRune(chars, thisLexer.next()) {
	}
	thisLexer.backup()
}

// errorf returns an error Token and halts the lexer
func (thisLexer *lexer) errorf(format string, args ...interface{}) stateFunctor {
	thisLexer.tokens <- Token{
		Type:  TokenError,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

// current returns the current string slice
func (thisLexer *lexer) current() string {
	return thisLexer.input[thisLexer.start:thisLexer.position]
}

// skip ignores all runes that match a string if no move has been done yet
func (thisLexer *lexer) skip(chars string) {
	// we can't skip if we moved already
	if thisLexer.hasMoved() {
		return
	}

	thisLexer.consume(chars)
	thisLexer.ignore()
}
