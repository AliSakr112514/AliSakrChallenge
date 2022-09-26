package TransactionService

import (
	"Challenge/config"
	Db "Challenge/internal/adapters/db"
	transacions "Challenge/internal/repositories/transactions"
	"github.com/google/uuid"
	"testing"
	"time"
)

func MainHandleConfiguration() {
	config := config.LoadConfig()
	var db, _ = Db.NewDatabaseConnection(config)
	transacions.Db = db
}
func TestCreateTransaction(t *testing.T) {
	MainHandleConfiguration()
	var trans transacions.Transaction
	trans.Id, _ = uuid.NewUUID()
	trans.Amount = 50
	trans.Isactive = false
	trans.Createdat = time.Now().String()
	trans.Currency = "GBP"
	result, error := TransactionService.CreateTransactionService(&trans)
	if result.Isactive != true || error != nil {
		t.Errorf("The test failed during creating transaction: %v", error)
	}
}
