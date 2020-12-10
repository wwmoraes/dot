package tokenizer

import (
	"strconv"
)

type stateFunctor func(Lexer) stateFunctor

func stateConsumeKeywordWith(keyword string, token TokenType) stateFunctor {
	return func(lexer Lexer) stateFunctor {
		lexer.advance(len(keyword))
		lexer.emit(token)

		return stateRoot
	}
}

func stateRoot(lexer Lexer) stateFunctor {
	for {
		lexer.skip(" \t\r\n")

		// check line comment
		if lexer.hasPrefix(KeywordLineComment) {
			if lexer.hasMoved() {
				// ignore anything before
				lexer.ignore()
			}
			return stateLineComment
		}
		// check strict flag
		if lexer.hasPrefix(KeywordStrict) {
			if lexer.hasMoved() {
				// ignore anything before
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordStrict, TokenStrict)
		}
		// check subgraph type
		if lexer.hasPrefix(KeywordSubgraph) {
			if lexer.hasMoved() {
				// ignore anything before
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordSubgraph, TokenSubgraph)
		}
		// check digraph type
		if lexer.hasPrefix(KeywordDigraph) {
			if lexer.hasMoved() {
				// ignore anything before
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordDigraph, TokenDigraph)
		}
		// check graph type
		if lexer.hasPrefix(KeywordGraph) {
			if lexer.hasMoved() {
				// ignore anything before
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordGraph, TokenGraph)
		}
		// check for start of block
		if lexer.hasPrefix(KeywordOpenBlock) {
			if lexer.hasMoved() {
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordOpenBlock, TokenOpenBlock)
		}
		// check for end of block
		if lexer.hasPrefix(KeywordCloseBlock) {
			if lexer.hasMoved() {
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordCloseBlock, TokenCloseBlock)
		}
		// check for start of attributes block
		if lexer.hasPrefix(KeywordOpenSquareBlock) {
			if lexer.hasMoved() {
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordOpenSquareBlock, TokenOpenSquareBlock)
		}
		// check for close of attributes block
		if lexer.hasPrefix(KeywordCloseSquareBlock) {
			if lexer.hasMoved() {
				lexer.ignore()
			}

			return stateConsumeKeywordWith(KeywordCloseSquareBlock, TokenCloseSquareBlock)
		}
		// check for semicolon
		if lexer.hasPrefix(KeywordSemicolon) {
			if lexer.hasMoved() {
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordSemicolon, TokenSemicolon)
		}
		// check for directed edge notation
		if lexer.hasPrefix(KeywordDirectedEdge) {
			if lexer.hasMoved() {
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordDirectedEdge, TokenDirectedEdge)
		}
		// check for undirected edge notation
		if lexer.hasPrefix(KeywordUndirectedEdge) {
			if lexer.hasMoved() {
				lexer.ignore()
			}
			return stateConsumeKeywordWith(KeywordUndirectedEdge, TokenUndirectedEdge)
		}
		// check for invalid angled brackets
		if lexer.hasPrefix(">") {
			return lexer.errorf("unexpected >")
		}
		//check for attributes
		for _, keywordAttribute := range KeywordAttributes {
			if lexer.hasPrefix(keywordAttribute + "=") {
				if lexer.hasMoved() {
					lexer.ignore()
				}

				return stateAttributeWith(keywordAttribute)
			}
		}
		// string is probably an identifier, so parse it as such
		if lexer.hasMoved() {
			lexer.backup()
			return stateIdentifier
		}

		if lexer.next() == RuneEOF {
			break
		}
	}

	// return to the last emit point and consume "ok-ish" characters
	lexer.restart()
	lexer.skip(" \t\r\n")

	// ignore anything not identified until EOF
	if lexer.hasMoved() {
		return lexer.errorf("unknown token '%s'", lexer.current())
	}

	lexer.emit(TokenEOF)
	return nil
}

func stateLineComment(lexer Lexer) stateFunctor {
	lexer.advance(len(KeywordLineComment))

	for {
		if lexer.next() == RuneEOF {
			break
		}

		if lexer.hasPrefix("\r\n") {
			break
		}

		if lexer.hasPrefix("\n") {
			break
		}

		if lexer.hasPrefix("\r") {
			break
		}
	}

	lexer.emit(TokenLineComment)

	return stateRoot
}

func stateIdentifier(lexer Lexer) stateFunctor {
	// parse a quoted identifier
	if lexer.hasPrefix(`"`) {
		return stateQuotedIdentifier
	}
	// parse an HTML identifier
	if lexer.hasPrefix(`<`) {
		return stateHTMLIdentifier
	}

	for {
		if lexer.hasPrefix(" ") {
			break
		}

		if lexer.hasPrefix("\n") {
			break
		}

		if lexer.hasPrefix("\r") {
			break
		}

		if lexer.hasPrefix(`"`) {
			break
		}

		if lexer.hasPrefix(`;`) {
			break
		}

		if lexer.hasPrefix(`,`) {
			break
		}

		// error if a non-printable, control character is found
		if !strconv.IsPrint(lexer.peek()) {
			return lexer.errorf("invalid characters on identifier '%s'", lexer.current())
		}

		if lexer.next() == RuneEOF {
			break
		}
	}

	lexer.emit(TokenLiteralID)

	return stateRoot
}

func stateQuotedIdentifier(lexer Lexer) stateFunctor {
	// accept the open quote we are on currently
	lexer.accept(`"`)
	for {
		// accept escaped quotes
		if lexer.hasPrefix(`\"`) {
			lexer.advance(2)
			continue
		}
		// accept newlines
		if lexer.accept("\r\n") {
			continue
		}
		// accept closing quote
		if lexer.accept(`"`) {
			break
		}
		// error if a non-printable, control character is found
		if !strconv.IsPrint(lexer.peek()) {
			return lexer.errorf("invalid characters on identifier '%s'", lexer.current())
		}
		// error if EOF is found before the quotes are closed
		if lexer.next() == RuneEOF {
			return lexer.errorf("unclosed identifier '%s'", lexer.current())
		}
	}

	lexer.emit(TokenQuotedID)

	return stateRoot
}

func stateHTMLIdentifier(lexer Lexer) stateFunctor {
	// accept the open angled bracket we are on currently
	lexer.accept("<")
	var pairs = 1
	for {
		// break loop if all angled brackets pairs have matched
		if pairs == 0 {
			break
		}
		// accept tag/this identifier close angled bracket
		if lexer.accept(">") {
			pairs--
			continue
		}
		// accept tag/this identifier open angled bracket
		if lexer.accept("<") {
			pairs++
			continue
		}
		// get next char, and error if EOF is found before the identifier closure
		if lexer.next() == RuneEOF {
			return lexer.errorf("unclosed identifier '%s'", lexer.current())
		}
	}

	lexer.emit(TokenHTMLID)

	return stateRoot
}

func stateAttributeWith(keyword string) stateFunctor {
	return func(lexer Lexer) stateFunctor {
		lexer.advance(len(keyword))
		lexer.emit(TokenAttribute)

		if !lexer.hasPrefix("=") {
			return lexer.errorf("expected = and value after attribute %s", keyword)
		}

		lexer.advance(len(KeywordEquals))
		lexer.emit(TokenEquals)

		return stateIdentifier
	}
}
