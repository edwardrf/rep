package main

import (
	"compress/flate"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	t := []byte(`package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
	"os"
)

func main() {
	t := []byte(X)
	p(t, os.Stdout, true)
}

func p(t []byte, w io.Writer, o bool) {
	q := []byte{96}
	var r io.Reader
	r = bytes.NewReader(t)
	if o {
		r = flate.NewReader(base64.NewDecoder(base64.RawStdEncoding, r))
	}
	s := bufio.NewReader(r)
	for {
		b, err := s.ReadByte()
		if err != nil {
			break
		}
		if o && b == 88 {
			w.Write(q)
			p(t, w, false)
			w.Write(q)
		} else {
			w.Write([]byte{b})
		}
	}
}
`)

	for i := 0; i <= 9; i++ {
		fmt.Printf("============\n[")
		//o := lzw.NewWriter(ascii85.NewEncoder(os.Stdout), lzw.LSB, i)
		//o := lzw.NewWriter(base64.NewEncoder(base64.RawStdEncoding, os.Stdout), lzw.LSB, i)
		o, err := flate.NewWriter(base64.NewEncoder(base64.RawStdEncoding, os.Stdout), i)
		if err != nil {
			fmt.Printf("NERR: %v", err)
		}
		_, err = o.Write(t)
		if err != nil {
			fmt.Printf("ERR: %v", err)
		}
		o.Close()
		fmt.Printf("]\n============\n")
	}
}
