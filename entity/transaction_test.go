package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmouthGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransactionWithAmouthLesserThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater that 1", err.Error())
}
