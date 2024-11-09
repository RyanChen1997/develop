package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filename := "text"

	// 读取文件内容
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 计算文件的原始大小
	originalSize := len(data)
	fmt.Printf("Original Size: %d bytes\n", originalSize)

	// 对内容进行压缩
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		log.Fatalf("Failed to write to gzip writer: %v", err)
	}
	if err := gz.Close(); err != nil {
		log.Fatalf("Failed to close gzip writer: %v", err)
	}

	// 计算压缩后的大小
	compressedSize := buf.Len()
	fmt.Printf("Compressed Size: %d bytes\n", compressedSize)
}
