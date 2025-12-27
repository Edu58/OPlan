package repository

import (
	"context"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountTypeRepository struct {
	db *pgxpool.Pool
}

func NewAccountTypeRepository(db *pgxpool.Pool) *AccountTypeRepository {
	return &AccountTypeRepository{db}
}

func (a *AccountTypeRepository) Create(ctx context.Context, arg domain.CreateAccountTypeParams) (*domain.AccountType, error) {
	if err := arg.Validate(); err != nil {
		return nil, err
	}

	query := `
	INSERT INTO account_types
		(name, active)
		VALUES($1, $2)
	RETURNING id, name, active, inserted_at, updated_at
	`

	row := a.db.QueryRow(ctx, query, arg.Name, arg.Active)

	var i domain.AccountType

	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Active,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (a *AccountTypeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM account_types
	WHERE id = $1
	RETURNING id, name, active, inserted_at, updated_at
	`
	result, err := a.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return domain.ErrAccountTypeNotFound
	}

	return nil
}

func (a *AccountTypeRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.AccountType, error) {
	query := `
	SELECT id, name, active, inserted_at, updated_at FROM account_types
	WHERE id = $1 LIMIT 1
	`
	row := a.db.QueryRow(ctx, query, id)

	var i domain.AccountType

	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Active,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (a *AccountTypeRepository) GetByName(ctx context.Context, name string) (*domain.AccountType, error) {
	query := `
	SELECT id, name, active, inserted_at, updated_at FROM account_types
	WHERE name = $1 LIMIT 1
	`

	row := a.db.QueryRow(ctx, query, name)

	var i domain.AccountType
	err := row.Scan(

		&i.ID,
		&i.Name,
		&i.Active,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (a *AccountTypeRepository) List(ctx context.Context) ([]*domain.AccountType, error) {
	query := `
	SELECT id, name, active, inserted_at, updated_at FROM account_types
	ORDER BY inserted_at DESC
	`
	rows, err := a.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []*domain.AccountType{}

	for rows.Next() {
		var i domain.AccountType
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Active,
			&i.InsertedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, &i)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (a *AccountTypeRepository) Update(ctx context.Context, arg domain.UpdateAccountTypeByIDParams) (*domain.AccountType, error) {
	query := `
	UPDATE account_types
	SET name=$2, active=$3
	WHERE id = $1
	RETURNING id, name, active, inserted_at, updated_at
	`

	row := a.db.QueryRow(ctx, query, arg.ID, arg.Name, arg.Active)
	var i domain.AccountType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Active,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (a *AccountTypeRepository) Exists(ctx context.Context, name string) (bool, error) {
	query := `
	SELECT EXISTS (
	    SELECT 1
	    FROM account_types
	    WHERE name = $1
	)
	`

	row := a.db.QueryRow(ctx, query, name)

	var exists bool
	err := row.Scan(&exists)

	return exists, err
}
