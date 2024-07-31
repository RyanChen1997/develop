package main

import (
	"bytes"
	"io"
	"log"
)

func TestIO() {
	var s = "hello world"
	r := bytes.NewReader([]byte(s))

	var buf = make([]byte, 0, 10)
	n, err := r.Read(buf)
	if err != nil {
		panic(err)
	}
	log.Printf("read %d bytes, %s\n", n, string(buf))

	io.ReadAll(r)
}
