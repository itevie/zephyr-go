package errors

type ErrorType uint

const (
	LexerError   = iota
	ParserError  = iota
	RuntimeError = iota
)

type Error struct {
	Message   string
	ErrorType ErrorType
}

func (error Error) Error() string {
	return error.Message
}
