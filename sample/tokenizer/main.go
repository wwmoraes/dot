package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/wwmoraes/dot/tokenizer"
)

func main() {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if _, file, _, ok := runtime.Caller(0); ok {
		dirname = path.Dir(file)
	}

	fileBytes, err := ioutil.ReadFile(path.Join(dirname, "test.dot"))
	if err != nil {
		log.Fatal(err)
	}

	graphString := string(fileBytes)
	log.Printf("graph contents:\n%s\n", graphString)

	log.Println("starting lexer...")
	_, tokenChan := tokenizer.NewRunnerLexer("sample", graphString)

	log.Println("opening plain.dot file to write...")
	fdPlain, err := os.Create("plain.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer fdPlain.Close()
	log.Println("opening pretty.dot file to write...")
	fdPretty, err := os.Create("pretty.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer fdPretty.Close()

	log.Println("consuming lexer tokens...")
	done := make(chan bool)
	go func() {
		type RuleMap map[tokenizer.TokenType]struct{}

		type TokenRule struct {
			Printer     func(w io.Writer, a ...interface{}) (n int, err error)
			SpaceAfter  RuleMap
			IndentAfter RuleMap
		}

		rules := map[tokenizer.TokenType]TokenRule{
			tokenizer.TokenLiteralID: {
				Printer: fmt.Fprint,
				SpaceAfter: RuleMap{
					tokenizer.TokenGraph:    {},
					tokenizer.TokenSubgraph: {},
				},
				IndentAfter: RuleMap{
					tokenizer.TokenOpenBlock:  {},
					tokenizer.TokenCloseBlock: {},
					tokenizer.TokenSemicolon:  {},
				},
			},
			tokenizer.TokenQuotedID: {
				Printer: fmt.Fprint,
				SpaceAfter: RuleMap{
					tokenizer.TokenGraph:    {},
					tokenizer.TokenSubgraph: {},
				},
				IndentAfter: RuleMap{
					tokenizer.TokenOpenBlock:  {},
					tokenizer.TokenCloseBlock: {},
					tokenizer.TokenSemicolon:  {},
				},
			},
			tokenizer.TokenAttribute: {
				IndentAfter: RuleMap{
					tokenizer.TokenOpenBlock:  {},
					tokenizer.TokenCloseBlock: {},
				},
			},
			tokenizer.TokenOpenBlock: {
				Printer: fmt.Fprintln,
				SpaceAfter: RuleMap{
					tokenizer.TokenGraph:     {},
					tokenizer.TokenDigraph:   {},
					tokenizer.TokenSubgraph:  {},
					tokenizer.TokenLiteralID: {},
					tokenizer.TokenQuotedID:  {},
				},
			},
			tokenizer.TokenOpenSquareBlock: {
				Printer: fmt.Fprint,
				SpaceAfter: RuleMap{
					tokenizer.TokenGraph: {},
				},
			},
			tokenizer.TokenCloseBlock: {
				Printer: fmt.Fprintln,
				IndentAfter: RuleMap{
					tokenizer.TokenOpenBlock:  {},
					tokenizer.TokenCloseBlock: {},
					tokenizer.TokenSemicolon:  {},
				},
			},
			tokenizer.TokenSemicolon: {
				Printer: fmt.Fprintln,
			},
			tokenizer.TokenSubgraph: {
				IndentAfter: RuleMap{
					tokenizer.TokenOpenBlock:  {},
					tokenizer.TokenCloseBlock: {},
					tokenizer.TokenSemicolon:  {},
				},
			},
			tokenizer.TokenGraph: {
				IndentAfter: RuleMap{
					tokenizer.TokenOpenBlock:  {},
					tokenizer.TokenCloseBlock: {},
					tokenizer.TokenSemicolon:  {},
				},
			},
		}

		level := 0
		previousTokenType := tokenizer.TokenEOF

		for {
			select {
			case token, open := <-tokenChan:
				if !open {
					done <- true
					return
				}

				log.Printf("%s[%++v] => %s\n", tokenizer.GetTokenName(token.Type), token.Type, token.Value)

				// a space is needed before unquoted identifiers
				if token.Type == tokenizer.TokenLiteralID {
					fmt.Fprint(fdPlain, " ")
				}

				// print to the plain file
				fmt.Fprint(fdPlain, token.Value)

				if token.Type == tokenizer.TokenCloseBlock {
					level--
				}

				// apply rules for pretty file
				if ruleSet, ok := rules[token.Type]; ok {
					// check if the space after applies
					if _, ok := ruleSet.SpaceAfter[previousTokenType]; ok {
						fmt.Fprint(fdPretty, " ")
					}
					// indent after
					if _, ok := ruleSet.IndentAfter[previousTokenType]; ok {
						fmt.Fprint(fdPretty, strings.Repeat("  ", level))
					}
					if ruleSet.Printer != nil {
						ruleSet.Printer(fdPretty, token.Value)
					} else {
						fmt.Fprint(fdPretty, token.Value)
					}
				} else {
					fmt.Fprint(fdPretty, token.Value)
				}

				if token.Type == tokenizer.TokenOpenBlock {
					level++
				}

				// format pretty file
				// switch token.Type {
				// case tokenizer.TokenLiteralID:
				// 	rule := map[tokenizer.TokenType]struct{}{
				// 		tokenizer.TokenGraph:    {},
				// 		tokenizer.TokenSubgraph: {},
				// 	}
				// 	if _, ok := rule[previousTokenType]; ok {
				// 		fmt.Fprint(fdPretty, " ")
				// 	}
				// 	fmt.Fprint(fdPretty, token.Value)
				// case tokenizer.TokenQuotedID:
				// 	fmt.Fprint(fdPretty, " ")
				// 	fmt.Fprint(fdPretty, token.Value)
				// case tokenizer.TokenOpenBlock:
				// 	fmt.Fprint(fdPretty, " ")
				// 	fmt.Fprintln(fdPretty, token.Value)
				// 	level++
				// 	fmt.Fprint(fdPretty, strings.Repeat("  ", level))
				// case tokenizer.TokenCloseBlock:
				// 	level--
				// 	fmt.Fprintln(fdPretty, token.Value)
				// 	fmt.Fprint(fdPretty, strings.Repeat("  ", level))
				// case tokenizer.TokenSemicolon:
				// 	fmt.Fprintln(fdPretty, token.Value)
				// 	fmt.Fprint(fdPretty, strings.Repeat("  ", level))
				// default:
				// 	fmt.Fprint(fdPretty, token.Value)
				// }

				previousTokenType = token.Type
				break
			case <-time.After(time.Second * 600):
				log.Println("lexer channel timed out")
				close(tokenChan)
				break
			}
		}
	}()
	<-done

	log.Println("done!")
}
