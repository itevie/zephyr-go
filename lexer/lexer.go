package lexer

import (
	"regexp"
	"strings"
	"zephyr-go/errors"
)

type Keyword struct {
	Matches   string
	TokenType TokenTypeType
}

func Lex(input string) ([]Token, error) {
	// Construct keywords
	var keywords = []Keyword{
		{Matches: "var", TokenType: TT_VariableDeclaration},
	}

	// Construct operators
	var operators = []Keyword{
		{Matches: "=", TokenType: TT_Assignment},
	}

	var tokens = []Token{}

	var chars = strings.Split(input, "")

	var identifierRegex = regexp.MustCompile("[a-z_]")
	var numberRegex = regexp.MustCompile("[0-9]")
	var uselessWhitespaceRegex = regexp.MustCompile("[ \r]")

	for len(chars) != 0 {
		var tokenValue string = ""
		var tokenType TokenTypeType = TT_Unknown

		// Used to remove the current character and return it
		var eat = func() string {
			var old = chars[0]
			chars = chars[1:]
			return old
		}

		// Used to set the value
		var setToken = func(value string, t TokenTypeType) {
			tokenValue = value
			tokenType = t
		}

		if uselessWhitespaceRegex.MatchString(chars[0]) { // WHITESPACE
			eat()
			continue
		} else if identifierRegex.MatchString(chars[0]) { // IDENTIFIER
			var identifier = eat()

			// Repeat until there is no more
			for len(chars) > 0 && identifierRegex.MatchString(chars[0]) {
				identifier += eat()
			}

			// Check if it is a keyword
			var alreadySet = false
			for _, element := range keywords {
				if element.Matches == identifier {
					setToken(identifier, element.TokenType)
					alreadySet = true
				}
			}

			if !alreadySet {
				setToken(identifier, TT_Identifier)
			}
		} else if numberRegex.MatchString(chars[0]) { // NUMBER
			var number = eat()

			for len(chars) > 0 && numberRegex.MatchString(chars[0]) {
				number += eat()
			}

			setToken(number, TT_Number)
		} else { // UNKNOWN
			// Check if it is an operator
			var alreadySet = false
			for _, element := range operators {
				if element.Matches == chars[0] {
					setToken(eat(), element.TokenType)
					alreadySet = true
				}
			}
			if !alreadySet {
				// Unknown thing given
				return tokens, &errors.Error{
					Message:   "Unknown character: " + chars[0],
					ErrorType: errors.LexerError,
				}
			}
		}

		// Add the token
		tokens = append(tokens, Token{
			Value:     tokenValue,
			TokenType: tokenType,
		})
	}

	return tokens, nil
}
