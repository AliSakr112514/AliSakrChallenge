package api

import (
	TransactionRepo "Challenge/internal/repositories/transactions"
	TransactionService "Challenge/internal/services/transactionsService"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

//Get all transactions
func GetAll(w http.ResponseWriter, r *http.Request) {

	transactions, err := TransactionRepo.TransactionRepo.GetAllTransactions()
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(200)
}

//create new transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	trans := TransactionRepo.Transaction{}
	json.NewDecoder(r.Body).Decode(&trans)
	_, err := TransactionService.TransactionService.CreateTransactionService(&trans)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't Add Transaction %s", err), 200)
		return
	}
	w.Write([]byte("The transaction added successfully"))
	w.WriteHeader(200)
}

func HandleRequest() {
	Router := chi.NewRouter()
	Router.Get("/transactions", GetAll)
	Router.Post("/transactions/create", CreateTransaction)
	log.Fatal(http.ListenAndServe(":8080", Router))
}
