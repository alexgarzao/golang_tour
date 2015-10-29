package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13.r.Read(b)
	for index := 0; index < n; index++ {
		if (b[index] >= 'A' && b[index] < 'N') || (b[index] >='a' && b[index] < 'n') {
			b[index] += 13
		} else if (b[index] > 'M' && b[index] <= 'Z') || (b[index] > 'm' && b[index] <= 'z'){
			b[index] -= 13
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
