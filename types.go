package gownload

type Options struct {
	Name   string
	Folder []string
}

type Result struct {
	Error error
	Path  string
}
