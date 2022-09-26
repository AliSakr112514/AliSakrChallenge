package transactions

import (
	log "Challenge/internal/adapters/Logger"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

type DefaultTransactionRepo struct {
	transaction ITransaction
}

var TransactionRepo DefaultTransactionRepo

func InitializeTransacionRepo(DatabasInitialization *gorm.DB) {
	Db = DatabasInitialization
}

func (s *DefaultTransactionRepo) GetAllTransactions() ([]Transaction, error) {
	var trans = []Transaction{}
	Db.Find(&trans)
	if trans != nil {
		log.Info.Printf("c31f9477-7360-4668-a70a-b2e3dfbef262 Get transactions Scucessfully, %v", trans)
		return trans, nil
	} else {
		log.Error.Printf("d4cd8940-5296-415b-a0e6-257d68709e9f Error occur while getting all transactions %v", trans)
		return nil, errors.New("no transactions found")
	}
}

func (s *DefaultTransactionRepo) CreateTransaction(trans *Transaction) (Transaction, error) {
	trans.Id = uuid.New()
	trans.Createdat = time.Now().String()
	trans.Isactive = false
	if err := Db.Create(&trans).Error; err != nil {
		return *trans, err
	}
	return *trans, nil
}
func (s *DefaultTransactionRepo) UpdateTransaction(trans *Transaction) (bool, error) {
	trans.Isactive = true
	if err := Db.Model(&trans).Update("isactive", true).Error; err != nil {
		return false, err
	}
	return true, nil
}
