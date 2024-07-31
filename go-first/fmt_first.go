package main

import (
	"fmt"
	"net/http"
)

type FmtStruct struct {
	Id    int
	Value string
}

func TestFmt() {
	s := FmtStruct{Id: 1, Value: "test"}
	fmt.Printf("%#v\n", s)

	err := http.ErrAbortHandler
	errNew := fmt.Errorf("WHO AM I %w", err)
	fmt.Println(errNew)
}
