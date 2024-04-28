package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

var re = regexp.MustCompile("(?i)^Error")

func commandFFmpeg() {
	convertWavToMp3("3.wav", "0.mp3")
	// commandRun()
}

func convertWavToMp3(origin, object string) error {
	cmd := exec.Command("ffmpeg", "-i", origin, "-f", "mp3", object)
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		text := scanner.Text()
		if containsError(text) {
			fmt.Println(text)
		}
	}

	return nil
}

// containsError 正则匹配text是否有Error开头
func containsError(text string) bool {
	return re.MatchString(text)
}

func commandRun() {
	cmd := exec.Command("ffmpeg", "-i", "1.wav", "-f", "mp3", "0.mp3")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
