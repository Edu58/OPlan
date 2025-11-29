package repository

import (
	"context"
	"testing"

	mock_db "github.com/Edu58/Oplan/internal/database/mock"
	db "github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewAccountTypeRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQueries := mock_db.NewMockQuerier(ctrl)

	ctx := context.Background()

	params := db.CreateAccountTypeParams{Name: "Test Acc Type", Active: pgtype.Bool{Bool: true}}
	expected := &db.AccountType{Name: "Test Acc Type", Active: pgtype.Bool{Bool: true}}

	mockQueries.EXPECT().
		CreateAccountType(gomock.Any(), gomock.AssignableToTypeOf(db.CreateAccountTypeParams{})).
		Return(expected, nil).
		Times(1)

	acc_type, err := NewAccountTypeRepository(mockQueries).CreateAccountType(ctx, params)

	assert.NoError(t, err)
	assert.IsType(t, &db.AccountType{}, acc_type)
	assert.Equal(t, acc_type.Name, expected.Name)
	assert.True(t, acc_type.Active.Bool)
}
