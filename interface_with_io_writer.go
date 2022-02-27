package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	writer io.Writer
}

func (c *Capper) Write(p []byte) (n int, e error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))

	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}

	return c.writer.Write(out)
}

func main() {

	c := &Capper{
		os.Stdout,
	}

	fmt.Fprintln(c, "Hey ! This is Gunjan's practice program in Go lang")
}
