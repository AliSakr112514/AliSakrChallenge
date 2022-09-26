package TransactionService

import (
	stream "Challenge/internal/adapters/stream"
	TransModel "Challenge/internal/repositories/transactions"
	"errors"
)

type DefaultTransactionService struct {
	transactionService ITransactionService
}

var TransactionService DefaultTransactionService

func (s *DefaultTransactionService) CreateTransactionService(trans *TransModel.Transaction) (TransModel.Transaction, error) {
	if trans.Amount < LowAmount || trans.Amount > HighAmount {
		return TransModel.Transaction{}, errors.New("The amount must be between 0 and 100000")
	}
	if len(trans.Currency) > NumberOfAvailableCharacter {
		return TransModel.Transaction{}, errors.New("Currency iso should be 3 char")
	}
	result, err := TransModel.TransactionRepo.CreateTransaction(trans)
	stream.Produce(trans)
	return result, err
}
