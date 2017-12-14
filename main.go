package main

import (
	"io"
	"os"
)

func main() {
	t := []byte(`package main

import (
	"io"
	"os"
)

func main() {
	t := []byte(X)
	p(t, os.Stdout, 1)
}
func p(t []byte, w io.Writer, o int) {
	q := []byte{96}
	for _, b := range t {
		if o > 0 && b == 88 {
			w.Write(q)
			p(t, w, 0)
			w.Write(q)
		} else {
			w.Write([]byte{b})
		}
	}
}
`)
	p(t, os.Stdout, 1)
}
func p(t []byte, w io.Writer, o int) {
	q := []byte{96}
	for _, b := range t {
		if o > 0 && b == 88 {
			w.Write(q)
			p(t, w, 0)
			w.Write(q)
		} else {
			w.Write([]byte{b})
		}
	}
}
