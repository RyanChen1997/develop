package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Shopify/sarama"
)

type VnTicketVoiceMsg struct {
	CallId        int64  `json:"callId"`
	FileName      string `json:"fileName"`
	CallDirection int    `json:"callDirection"`
	Callee        string `json:"callee"`
	Remark        string `json:"remark"` // driver_id 备注内容是司机的id
	Source        string `json:"source"`
	CityId        int    `json:"cityId"`
	Type          string `json:"type"`
	Caller        string `json:"caller"`
	Supplier      string `json:"supplier"`
	BizKey        int    `json:"bizKey"`
	BizId         string `json:"bizId"` // order_id 订单号
	SessionUuid   int64  `json:"sessionUuid"`
	FileUrl       string `json:"fileUrl"`
	Day           string `json:"day"`
	RecordMode    string `json:"recordMode"`
}

func main() {
	// 配置项
	brokers := []string{"kfk1-stg.myhll.cn:9092", "kfk2-stg.myhll.cn:9092", "kfk3-stg.myhll.cn:9092"} // Kafka broker地址
	topic := "vn_ticket_voice"                                                                        // 主题
	cityIds := []int{1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010}
	messageCount := 5
	messageInterval := 1500 * time.Millisecond // 发送间隔
	bizKey := 8

	// 设置生产者配置
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < messageCount; i++ {
		cityId := cityIds[i%len(cityIds)]
		msg := &VnTicketVoiceMsg{
			CallId:        rand.Int63(),
			FileName:      "2024/9/26/025e60de-f3d2-4aa2-bc44-2d025d2ca888.wav",
			CallDirection: 1,
			Callee:        "*************",
			Remark:        randID(),
			Source:        "bme-trade-biz-extra-core-svc",
			CityId:        cityId,
			Type:          "ticket_voice",
			Caller:        "***********",
			Supplier:      "10006",
			BizKey:        bizKey,
			BizId:         randID(),
			SessionUuid:   rand.Int63(),
			FileUrl:       "https://intelligent-operation-oss.oss-cn-shenzhen.aliyuncs.com/test/ryanchen/app_id_2_3.wav?OSSAccessKeyId=LTAI5t6zHU568fM6wrs3drRj&Expires=1730810407&Signature=RXdhzZjIAi90mLstq6J%2Bk2FJ%2F8g%3D",
			Day:           "20240927",
			RecordMode:    "1000",
		}

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Fatalln("Failed to marshal message:", err)
		}

		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(msgBytes),
		}

		_, _, err = producer.SendMessage(message)
		if err != nil {
			log.Println("Failed to send message:", err)
		} else {
			log.Println("Message sent:", msg)
		}

		time.Sleep(messageInterval)
	}
}

func randID() string {
	return fmt.Sprintf("%d", rand.Int63())
}
