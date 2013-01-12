package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rdr *rot13Reader) Read(p []byte) (n int, err error) {
	n,e := rdr.r.Read(p)
	if e == nil {
		for i,ch := range p {
			switch {
			case  ch < 'A' || ch > 'z':
			case ch + 13 > 'z':
				p[i] = 'a' + 12 - ('z' - ch)
			default:
				p[i]+=13
			}
		}
	}
	return n,e
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
