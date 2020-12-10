package tokenizer

// IteratorLexer is implemented by lexers that analyzes and returns tokens on
// demand, through a next method
type IteratorLexer interface {
	Lexer
	NextToken() (Token, bool)
}

type iteratorLexer struct {
	lexer
	state stateFunctor
}

// NewIteratorLexer creates a new instance of the dot lexer as an iterator one
func NewIteratorLexer(name, input string) IteratorLexer {
	lex := &iteratorLexer{
		lexer: lexer{
			input:  input,
			tokens: make(chan Token, 2),
		},
		state: stateRoot,
	}

	return lex
}

func (thisLexer *iteratorLexer) NextToken() (Token, bool) {
	for {
		select {
		case token, closed := <-thisLexer.tokens:
			return token, closed
		default:
			if thisLexer.state == nil {
				close(thisLexer.tokens)
			} else {
				thisLexer.state = thisLexer.state(thisLexer)
			}
		}
	}
}
