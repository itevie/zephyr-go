package lexer

type TokenTypeType uint8

const (
	// Literals
	TT_Unknown    = iota
	TT_Identifier = iota
	TT_Number     = iota

	// Operators
	TT_Assignment = iota

	// Keywords
	TT_VariableDeclaration = iota
)

type Token struct {
	Value     string
	TokenType TokenTypeType
}
