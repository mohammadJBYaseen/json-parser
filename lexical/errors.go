package lexical

type (
	ParseError struct {
		err     error
		message string
	}
)

func (f ParseError) Error() string {
	return f.message + ": " + f.err.Error()
}
