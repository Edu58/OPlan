package domain

import (
	"testing"

	db "github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/stretchr/testify/assert"
)

func TestNewAccountType(t *testing.T) {
	acc_type := NewAccountType("test", true)
	assert.NotNil(t, acc_type)
	assert.IsType(t, &db.AccountType{}, acc_type)
	assert.Equal(t, acc_type.Name, "test")
	assert.True(t, acc_type.Active.Bool)
	assert.NotNil(t, acc_type.InsertedAt)
	assert.NotNil(t, acc_type.UpdatedAt)
}
