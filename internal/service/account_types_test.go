package service

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/Edu58/Oplan/internal/domain"
	mock_repo "github.com/Edu58/Oplan/internal/repository/mocks"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountTypes(t *testing.T) {
	tests := []struct {
		name     string
		params   domain.CreateAccountTypeParams
		expected *domain.AccountType
		setup    func(mockRepo *mock_repo.MockAccountTypeRepository, params domain.CreateAccountTypeParams)
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "Create Active Account Type",
			params:   domain.CreateAccountTypeParams{Name: "Test Acc Type", Active: true},
			expected: &domain.AccountType{Name: "Test Acc Type", Active: true},
			setup: func(mockRepo *mock_repo.MockAccountTypeRepository, params domain.CreateAccountTypeParams) {
				mockRepo.EXPECT().
					Create(gomock.Any(), params).
					Return(&domain.AccountType{
						Name:   "Test Acc Type",
						Active: true},
						nil).
					Times(1)
			},
			wantErr: false,
		},
		{
			name:     "Create Inactive Account Type",
			params:   domain.CreateAccountTypeParams{Name: "Test Acc Type 2", Active: false},
			expected: &domain.AccountType{Name: "Test Acc Type 2", Active: false},
			setup: func(mockRepo *mock_repo.MockAccountTypeRepository, params domain.CreateAccountTypeParams) {
				mockRepo.EXPECT().
					Create(gomock.Any(), params).
					Return(&domain.AccountType{
						Name:   "Test Acc Type 2",
						Active: false},
						nil).
					Times(1)
			},
			wantErr: false,
		},
		{
			name:     "Create Duplicate Account Type",
			params:   domain.CreateAccountTypeParams{Name: "Test Acc Type 2", Active: false},
			expected: nil,
			setup: func(mockRepo *mock_repo.MockAccountTypeRepository, params domain.CreateAccountTypeParams) {
				mockRepo.EXPECT().
					Create(gomock.Any(), params).
					Return(nil, errors.New("duplicate key violates unique constraint")).
					Times(1)
			},
			wantErr: true,
			errMsg:  "duplicate key violates unique constraint",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			mockRepo := mock_repo.NewMockAccountTypeRepository(ctrl)

			test.setup(mockRepo, test.params)

			var buf bytes.Buffer
			testLogger := logger.NewLogger(&buf)

			acc_type, err := NewAccountTypesService(mockRepo, testLogger).Create(ctx, test.params)

			if test.wantErr {
				assert.Error(t, err)
				assert.Nil(t, acc_type)
				assert.Contains(t, err.Error(), test.errMsg)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, test.expected, acc_type)
				assert.Equal(t, acc_type.Name, test.expected.Name)
				assert.Equal(t, acc_type.Active, test.expected.Active)
			}
		})
	}
}
