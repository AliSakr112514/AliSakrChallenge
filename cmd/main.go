package main

import (
	"Challenge/config"
	route "Challenge/internal/adapters/api"
	Db "Challenge/internal/adapters/db"
	Stream "Challenge/internal/adapters/stream"
	TransacionRepo "Challenge/internal/repositories/transactions"
)

func main() {
	config := config.LoadConfig()
	var db, _ = Db.NewDatabaseConnection(config)
	TransacionRepo.InitializeTransacionRepo(db)
	Stream.InitializeKafka(config)
	go Stream.Consume()
	route.HandleRequest()
}
