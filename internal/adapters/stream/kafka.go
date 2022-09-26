package stream

import (
	"Challenge/config"
	log "Challenge/internal/adapters/Logger"
	TransModel "Challenge/internal/repositories/transactions"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
	"time"
)

var Kafkaconfig = kafka.ReaderConfig{}

func InitializeKafka(configurations *config.Configurations) {
	Kafkaconfig.Brokers = configurations.Kafka.Brokers
	Kafkaconfig.Topic = configurations.Kafka.Topic
	Kafkaconfig.MaxBytes = configurations.Kafka.MaxBytes
}
func Produce(transaction *TransModel.Transaction) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_transaction", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
}
func Consume() {
	reader := kafka.NewReader(Kafkaconfig)
	var transaction TransModel.Transaction
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {
			log.Error.Println(time.Now().String()+":: Error happened during calling kafka server %v", error)
			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
		json.Unmarshal(message.Value, &transaction)
		TransModel.TransactionRepo.UpdateTransaction(&transaction)
	}
}
