package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// use aliyun oss sdk
// use copyObjectFrom function to copy object
// set option to record process

var (
	Endpoint        = ""
	AccessKeyID     = ""
	AccessKeySecret = ""
	BucketName      = ""
)

func main() {
	copyObject("cuda-10.2.tar.gz", "test/ryanchen/cuda-10.2.tar.gz")
}

func copyObject(objectName, targetName string) {
	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	res, err := bucket.CopyObjectFrom(BucketName, objectName, targetName, oss.Progress(&ProcessListener{}))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("CopyObject success", res)
}

type ProcessListener struct {
}

func (listener *ProcessListener) ProgressChanged(event *oss.ProgressEvent) {
	fmt.Printf("event: %+v\n", event)

	if event.EventType == oss.TransferStartedEvent {
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n", event.ConsumedBytes, event.TotalBytes)
	} else if event.EventType == oss.TransferDataEvent {
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, RwBytes: %d.", event.ConsumedBytes, event.TotalBytes, event.RwBytes)
	} else if event.EventType == oss.TransferCompletedEvent {
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n", event.ConsumedBytes, event.TotalBytes)
	} else if event.EventType == oss.TransferFailedEvent {
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n", event.ConsumedBytes, event.TotalBytes)
	}
}
