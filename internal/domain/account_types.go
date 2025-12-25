package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AccountType struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Active     bool      `json:"active"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateAccountTypeParams struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type UpdateAccountTypeByIDParams struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Active bool      `json:"active"`
}

// ListAccountTypesFilter for filtering/pagination
type ListAccountTypesFilter struct {
	ActiveOnly bool
	Limit      int
	Offset     int
}

func (r *CreateAccountTypeParams) Validate() error {
	return ValidateAccountTypeName(r.Name)
}

func (r *UpdateAccountTypeByIDParams) Validate() error {
	return ValidateAccountTypeName(r.Name)
}

func ValidateAccountTypeName(name string) error {
	if name == "" {
		return ErrInvalidAccountTypeName
	}
	if len(name) < 3 {
		return ErrAccountTypeNameTooShort
	}
	if len(name) > 50 {
		return ErrAccountTypeNameTooLong
	}
	return nil
}

type AccountTypeService interface {
	// Create a new account type with validation and business rules
	Create(ctx context.Context, req CreateAccountTypeParams) (*AccountType, error)

	// GetByID retrieves an account type by ID
	// GetByID(ctx context.Context, id string) (*AccountType, error)

	// // GetByName retrieves an account type by name
	GetByName(ctx context.Context, name string) (*AccountType, error)

	// // ListActive returns only active account types
	// ListActive(ctx context.Context) ([]*AccountType, error)

	// ListAll returns all account types
	ListAll(ctx context.Context) ([]*AccountType, error)

	// // Update modifies an existing account type
	// Update(ctx context.Context, req UpdateAccountTypeByIDParams) (*AccountType, error)

	// // Deactivate soft-deletes an account type
	// Deactivate(ctx context.Context, id string) error

	// // Activate re-activates a deactivated account type
	// Activate(ctx context.Context, id string) error

	// // Delete permanently removes an account type (use with caution)
	// Delete(ctx context.Context, id string) error
}

type AccountTypeRepository interface {
	Create(ctx context.Context, req CreateAccountTypeParams) (*AccountType, error)
	GetByID(ctx context.Context, id uuid.UUID) (*AccountType, error)
	GetByName(ctx context.Context, name string) (*AccountType, error)
	List(ctx context.Context) ([]*AccountType, error)
	Update(ctx context.Context, req UpdateAccountTypeByIDParams) (*AccountType, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, name string) (bool, error)
}
