package main

import (
	"os"

	"github.com/go-audio/wav"
)

func buildRecordAndConvert() {
	if len(os.Args) < 3 {
		panic("Boom.......")
	}

	mode := os.Args[1]

	switch mode {
	case "record":
		fileName := os.Args[2]
		recordMicrophoneToPCM(fileName)
	case "convert":
		if len(os.Args) < 4 {
			panic("Boom.......")
		}
		from := os.Args[2]
		to := os.Args[3]
		convertPCM2Wav(wav.Encoder{
			SampleRate:     16000,
			BitDepth:       16,
			NumChans:       1,
			WavAudioFormat: 1,
		}, from, to)
	default:
		panic("Boom.......")
	}
}
