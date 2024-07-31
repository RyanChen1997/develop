package main

import (
	"log"
	"os"
)

func main() {
	a, err := NewArgs("a#,b,c", os.Args)
	if err != nil {
		log.Print(err)
	}
	log.Printf("a: %v, b: %v, c: %v", a.GetInt("a"), a.GetBool("b"), a.GetBool("c"))
}
