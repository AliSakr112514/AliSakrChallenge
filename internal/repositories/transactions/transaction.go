package transactions

type ITransaction interface {
	CreateTransaction(trans *Transaction) (Transaction, error)
	GetAllTransactions() ([]Transaction, error)
	UpdateTransaction(JsonTransaction string) (bool, error)
}
