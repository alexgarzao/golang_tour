package main

import "golang.org/x/tour/reader"

type MyReader struct {
	Reader
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

func (mr MyReader) Read(b []byte) (int, error) {
	for index := range b {
		b[index] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
