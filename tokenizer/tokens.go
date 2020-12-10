package tokenizer

type Token struct {
	Type  TokenType
	Value string
}

type TokenType byte

const (
	TokenEOF              TokenType = iota
	TokenError            TokenType = iota
	TokenStrict           TokenType = iota
	TokenGraph            TokenType = iota
	TokenDigraph          TokenType = iota
	TokenSubgraph         TokenType = iota
	TokenLineComment      TokenType = iota
	TokenOpenBlock        TokenType = iota
	TokenCloseBlock       TokenType = iota
	TokenLiteralID        TokenType = iota
	TokenQuotedID         TokenType = iota
	TokenHTMLID           TokenType = iota
	TokenSemicolon        TokenType = iota
	TokenOpenSquareBlock  TokenType = iota
	TokenCloseSquareBlock TokenType = iota
	TokenDirectedEdge     TokenType = iota
	TokenUndirectedEdge   TokenType = iota
	TokenAttribute        TokenType = iota
	TokenEquals           TokenType = iota

	// GraphGlobalAttributes    TokenType = iota
	// NodeGlobalAttributes     TokenType = iota
	// EdgeGlobalAttributes     TokenType = iota
	// PortID                   TokenType = iota
	// PortCompass              TokenType = iota
	// PreprocessorInstruction  TokenType = iota
	// MultilineStringBackslash TokenType = iota

	// OpenSquareBracket  TokenType = iota
	// CloseSquareBracket TokenType = iota
	// OpenAngleBracket   TokenType = iota
	// CloseAngleBracket  TokenType = iota
	// OpenBlockComment   TokenType = iota
	// CloseBlockComment  TokenType = iota

	// DoubleQuote TokenType = iota
	// String      TokenType = iota
	// Plus        TokenType = iota
	// Colon       TokenType = iota
	// Comma       TokenType = iota
	// Space       TokenType = iota
)

var tokenName = map[TokenType]string{
	TokenEOF:              "EOF",
	TokenError:            "Error",
	TokenStrict:           "Strict",
	TokenGraph:            "Graph",
	TokenDigraph:          "Digraph",
	TokenSubgraph:         "Subgraph",
	TokenLineComment:      "LineComment",
	TokenOpenBlock:        "OpenBlock",
	TokenCloseBlock:       "CloseBlock",
	TokenLiteralID:        "LiteralID",
	TokenQuotedID:         "QuotedID",
	TokenHTMLID:           "HTMLID",
	TokenSemicolon:        "Semicolon",
	TokenOpenSquareBlock:  "OpenSquareBlock",
	TokenCloseSquareBlock: "CloseSquareBlock",
	TokenDirectedEdge:     "DirectedEdge",
	TokenUndirectedEdge:   "UndirectedEdge",
	TokenAttribute:        "Attribute",
	TokenEquals:           "Equals",
}

func GetTokenName(tokenType TokenType) string {
	name, ok := tokenName[tokenType]
	if !ok {
		name = "unknown"
	}

	return name
}
