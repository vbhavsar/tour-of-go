package main

import (
  "io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rdr *rot13Reader) Read(p []byte) (n int, err error) {
	n,e := rdr.r.Read(p)
	if e == nil {
		for i:=0; i<len(p); i++ {
			switch {
			case  p[i] < 'A' || p[i] > 'z':
			case p[i] + 13 > 'z':
				p[i] = 'a' + 12 - ('z' - p[i])
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
