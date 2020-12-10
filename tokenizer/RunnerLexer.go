package tokenizer

// RunnerLexer is implemented by lexers that run as goroutines and expose the
// tokens directly through a channel
type RunnerLexer interface {
	Lexer
	run()
}

// NewRunnerLexer creates a new instance of the dot lexer as a runner one
func NewRunnerLexer(name, input string) (RunnerLexer, chan Token) {
	lex := &lexer{
		input:  input,
		tokens: make(chan Token, 2),
	}

	go lex.run()

	return lex, lex.tokens
}

func (thisLexer *lexer) run() {
	for state := stateRoot; state != nil; {
		state = state(thisLexer)
	}
	close(thisLexer.tokens)
}
