package process_transaction

import (
	"testing"

	mock_entity "github.com/denis-souzaa/go-with-tdd/entity/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionWhenItsValid(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransactionWhenItsInvalid(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    2000,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "rejected",
		ErrorMessage: "you dont have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "rejected", "you dont have limit for this transaction").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
