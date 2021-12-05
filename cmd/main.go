package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/denis-souzaa/go-with-tdd/adapter/repository"
	"github.com/denis-souzaa/go-with-tdd/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "transactions.db")

	if err != nil {
		fmt.Println(err.Error())
	}

	repository := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repository)

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    0,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)

	fmt.Println(string(outputJson))
}
