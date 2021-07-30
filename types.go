package gownload

type Options struct {
	Name   string
	Folder []string
}

type Result struct {
	Error error
	Path  string
}

// Unpack & handle result easily.
// Will panic with given message if Result has an error,
// Otherwise it will run the given function.
func (result Result) Handle(msg error, fn func(path string)) {
	if result.Error != nil {
		panic(msg)
	} else {
		fn(result.Path)
	}
}
