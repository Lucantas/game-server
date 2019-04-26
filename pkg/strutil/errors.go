package strutil

// NotAstringErr is called always that one of the parameters
// sent to a function does not match a string or slice of bytes
type NotAstringErr struct {
	message string
}

func (err NotAstringErr) Error() string {
	return err.message
}

func newNotAstringErr(message string) error {
	return NotAstringErr{message}
}
