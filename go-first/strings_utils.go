package main

import (
	"log"
	"strings"
)

func lastIndex() {
	str := "abcdefg_ac_jb"
	ret := strings.LastIndex(str, "_")
	log.Println(ret)
}
