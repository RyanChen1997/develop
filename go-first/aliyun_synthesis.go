package main

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"
	"time"

	nls "github.com/aliyun/alibabacloud-nls-go-sdk"
)

const (
	//online key
	APPKEY = "" //获取Appkey请前往控制台：https://nls-portal.console.aliyun.com/applist
	TOKEN  = "" //获取Token具体操作，请参见：https://help.aliyun.com/document_detail/450514.html
)

type TtsUserParam struct {
	F      io.Writer
	Logger *nls.NlsLogger
}

func onTaskFailed(text string, param interface{}) {
	p, ok := param.(*TtsUserParam)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	p.Logger.Println("TaskFailed:", text)
}

func onSynthesisResult(data []byte, param interface{}) {
	p, ok := param.(*TtsUserParam)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}
	p.Logger.Println("收到语音数据，长度：", len(data))
	p.F.Write(data)
}

func onCompleted(text string, param interface{}) {
	p, ok := param.(*TtsUserParam)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	p.Logger.Println("onCompleted:", text)
}

func onClose(param interface{}) {
	p, ok := param.(*TtsUserParam)
	if !ok {
		log.Default().Fatal("invalid logger")
		return
	}

	p.Logger.Println("onClosed:")
}

func waitReady(ch chan bool, logger *nls.NlsLogger) error {
	select {
	case done := <-ch:
		{
			if !done {
				logger.Println("Wait failed")
				return errors.New("wait failed")
			}
			logger.Println("Wait done")
		}
	case <-time.After(60 * time.Second):
		{
			logger.Println("Wait timeout")
			return errors.New("wait timeout")
		}
	}
	return nil
}

var lk sync.Mutex
var fail = 0
var reqNum = 0

const (
	TEXT = "你好小德，今天天气怎么样。AI的潜力远远超过我们目前的理解和应用。随着技术的进步，AI将能够处理更复杂的任务，提供更深入的分析，甚至可能超越人类的决策能力。这种可能性引发了许多讨论和争议，但无论如何，AI的发展都是不可避免的。AI的发展将对社会产生深远影响。在经济方面，AI可以提高生产效率，降低运营成本，创造新的就业机会。在医疗领域，AI可以帮助医生进行诊断，提供个性化的治疗方案，甚至进行手术。在教育领域，AI可以提供个性化的学习体验，帮助教师进行教学。"
)

func testMultiInstance() {
	param := nls.DefaultSpeechSynthesisParam()
	config := nls.NewConnectionConfigWithToken(nls.DEFAULT_URL, APPKEY, TOKEN)
	strId := "ID0"
	fname := "ttsdump.wav"
	ttsUserParam := new(TtsUserParam)
	fout, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	logger := nls.NewNlsLogger(os.Stderr, strId, log.LstdFlags|log.Lmicroseconds)
	logger.SetLogSil(false)
	logger.SetDebug(true)
	logger.Printf("Test Normal Case for SpeechRecognition:%s", strId)
	ttsUserParam.F = fout
	ttsUserParam.Logger = logger
	//第三个参数控制是否请求长文本语音合成，false为短文本语音合成
	tts, err := nls.NewSpeechSynthesis(config, logger, false,
		onTaskFailed, onSynthesisResult, nil,
		onCompleted, onClose, ttsUserParam)
	if err != nil {
		logger.Fatalln(err)
		return
	}

	lk.Lock()
	reqNum++
	lk.Unlock()
	logger.Println("SR start")
	ch, err := tts.Start(TEXT, param, nil)
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		tts.Shutdown()
	}

	err = waitReady(ch, logger)
	if err != nil {
		lk.Lock()
		fail++
		lk.Unlock()
		tts.Shutdown()
	}
	logger.Println("Synthesis done")
	tts.Shutdown()

}
