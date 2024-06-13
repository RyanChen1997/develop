package main

import (
	"log"
	"strings"

	"github.com/antchfx/xmlquery"
)

const (
	ssml1 = `<speak pitch="-100">
	我的音高却比别人低。
  </speak>`
	ssml2 = `<speak rate="200" pitch="-100" volume="80">
  所以放在一起，我的声音是这样的。
</speak>`
	ssml3 = `<speak bgm="http://nls.alicdn.com/bgm/2.wav" backgroundMusicVolume="30" rate="-500" volume="40">
<break time="2s"/>
阴崖老木苍苍烟
<break time="700ms"/>
雨声犹在竹林间
<break time="700ms"/>
绵蕝固知裨国计
<break time="700ms"/>
绵州风物总堪怜
<break time="2s"/>
</speak>`
	ssml4 = `<speak>你好小德。</speak><speak>你好小德。</speak>`
)

func parseSSML() {
	doc1, err := xmlquery.Parse(strings.NewReader(ssml1))
	if err != nil {
		log.Fatal(err)
	}
	text1 := doc1.SelectElement("speak").InnerText()
	log.Println(text1)

	doc2, err := xmlquery.Parse(strings.NewReader(ssml2))
	if err != nil {
		log.Fatal(err)
	}
	speaks := doc2.SelectElements("speak")
	for _, speak := range speaks {
		text2 := speak.InnerText()
		log.Println(text2)
	}

	doc3, err := xmlquery.Parse(strings.NewReader(ssml3))
	if err != nil {
		log.Fatal(err)
	}
	text3 := doc3.SelectElement("speak").InnerText()
	log.Println(text3)

	doc4, err := xmlquery.Parse(strings.NewReader(ssml4))
	if err != nil {
		log.Fatal(err)
	}
	speaks = doc4.SelectElements("speak")
	for _, speak := range speaks {
		text4 := speak.InnerText()
		log.Println(text4)
	}
}
