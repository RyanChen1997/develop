package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// extractM3u8 启动一个Gin服务器，包含一个GET接口，用于提取m3u8文件
func extractM3u8() {
	os.Setenv("RUST_BACKTRACE", "1")

	// fmt.Println(catch("Detected https://c2.monidai.com/20230505/jWeMYiiP/index.m3u8"))
	// urls, err := extract("https://jjhanjutv.com/play/8331-0-4.html")
	// fmt.Println(err)
	// fmt.Println(urls)

	r := gin.Default()

	r.GET("/extract", func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "url is required",
			})
			return
		}

		fmt.Println("url:", url)

		urls, err := extract(url)
		if err != nil {
			fmt.Println("err:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Printf("request url: %s || response urls: %v\n", url, urls)
		c.JSON(http.StatusOK, gin.H{
			"urls": urls,
		})
	})

	r.Run(":8080")
}

func extract(url string) ([]string, error) {
	var urls []string
	// 创建一个exec.Cmd对象来执行命令
	cmd := exec.Command("vsd", "capture", url)

	// 创建一个定时器，当超过30秒时执行函数来杀死命令
	timeout := time.AfterFunc(10*time.Second, func() {
		err := cmd.Process.Signal(os.Interrupt)
		if err != nil {
			fmt.Printf("Error killing process: %s\n", err)
			return
		}
	})

	// 创建一个带有标准输出和错误输出的管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("Error creating StderrPipe: %v\n", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("Error creating StderrPipe: %v\n", err)
	}

	// 使用一个Scanner读取管道的输出
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			// fmt.Println(text)
			url, find := catch(text)
			if find {
				urls = append(urls, url)
			}
			timeout.Reset(10 * time.Second)
		}
	}()

	// 启动命令
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("Error starting command: %s", err.Error())
	}

	// 等待命令执行完毕
	err = cmd.Wait()
	if err != nil && !strings.Contains(err.Error(), "killed") {
		return nil, err
	}

	// 停止定时器
	timeout.Stop()

	// 关闭管道
	stderr.Close()
	return urls, nil
}

var pattern = regexp.MustCompile("^Detected (.*)$")

func catch(text string) (string, bool) {
	var (
		find bool
		url  string
	)

	matchs := pattern.FindStringSubmatch(text)
	if len(matchs) > 1 {
		find = true
		url = matchs[1]
	}
	return url, find
}
