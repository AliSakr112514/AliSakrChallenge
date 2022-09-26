package TransactionService

import TransModel "Challenge/internal/repositories/transactions"

type ITransactionService interface {
	CreateTransactionService(trans *TransModel.Transaction) (TransModel.Transaction, error)
}
