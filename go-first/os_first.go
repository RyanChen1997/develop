package main

import (
	"fmt"
	"os"
)

func TestOS() {
	// 定义一个映射，键是模板变量，值是替换的值
	mapping := map[string]string{
		"name": "World",
		"lang": "Go",
	}

	m := func(s string) string {
		return mapping[s]
	}

	// 使用模板变量的字符串
	templateStr := "Hello, ${name}. You are learning ${lang}."

	// 使用 os.Expand 替换模板变量
	result := os.Expand(templateStr, m)
	fmt.Println(result)
}
