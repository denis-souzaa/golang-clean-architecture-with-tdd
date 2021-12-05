package process_transaction

import (
	"github.com/denis-souzaa/go-with-tdd/entity"
)

type ProcessTransaction struct {
	Reposity entity.TransactionRepository
}

func NewProcessTransaction(reposity entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Reposity: reposity}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}

	return p.approveTransaction(transaction)
}

func (p *ProcessTransaction) approveTransaction(transaction *entity.Transaction) (TransactionDtoOutput, error) {
	err := p.Reposity.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "approved",
		ErrorMessage: "",
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidaTransaction error) (TransactionDtoOutput, error) {
	err := p.Reposity.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidaTransaction.Error())
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "rejected",
		ErrorMessage: invalidaTransaction.Error(),
	}

	return output, nil
}
