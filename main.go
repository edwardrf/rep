package main

import (
	"io"
	"os"
)

type B []byte

func main() {
	t := B(`package main

import (
	"io"
	"os"
)

type B []byte

func main() {
	t := B(X)
	p(t, os.Stdout, 1)
}
func p(t B, w io.Writer, o int) {
	for _, b := range t {
		if o > 0 && b == 88 {
			w.Write(B{96})
			p(t, w, 0)
			w.Write(B{96})
		} else {
			w.Write(B{b})
		}
	}
}
`)
	p(t, os.Stdout, 1)
}
func p(t B, w io.Writer, o int) {
	for _, b := range t {
		if o > 0 && b == 88 {
			w.Write(B{96})
			p(t, w, 0)
			w.Write(B{96})
		} else {
			w.Write(B{b})
		}
	}
}
