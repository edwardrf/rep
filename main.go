package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
	"os"
)

func main() {
	t := []byte(`ZJBBa+MwEIXPml8xm0OwwCSXJWQXfCnttYfk0ELpQbbHQcSxnJFcU4L+exnJ0JTexNObb+a90TRncyK8GDsA2MvoOGABalVPnXUreXwG8vJo3GVk8n7b9SaQKDQ0rrXDaVsbT7u/IuUZ51egAbppaBK50HgDFfB/hW/vAixeNaixCCU6vzmG1k2hxMATaYjL3FiExVzijNZtXtgG4hId1s71iXj9Jt7+7SKoD8PIYj6QaYlBMVaYEmyeac5iETQo26ETQjKkQHeGHEeER2rcnXIw8zG0T0vsEllrUBGUlztSY3cU1qA6x2lLXSIxi8un0x6kAw1K7pCPPxUOtk9WVTOZMyjh5jPXa6yxqnC/z4Y5V1FcBZBbnEvsTO9J//qPSL2nn4NLY3XUeU2ECF8BAAD/`)
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
