package repository

import (
	"context"

	"github.com/Edu58/Oplan/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Create(ctx context.Context, params domain.CreateUserParams) (*domain.User, error) {
	query := `
	INSERT INTO users
		(email, first_name, last_name, password, msisdn, email_verified, msisdn_verified, active, account_type_id)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING *
	`

	row := u.db.QueryRow(ctx, query, params.Email, params.FirstName, params.LastName, params.Password, params.MSISDN, params.EmailVerified, params.MsisdnVerified, params.Active, params.AccountTypeId)

	var i domain.User

	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.MSISDN,
		&i.EmailVerified,
		&i.MsisdnVerified,
		&i.Active,
		&i.AccountTypeId,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (u *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `
		SELECT * FROM users WHERE id = $1;
		`

	row := u.db.QueryRow(ctx, query, id)

	var i domain.User

	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.MSISDN,
		&i.EmailVerified,
		&i.MsisdnVerified,
		&i.Active,
		&i.AccountTypeId,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT * FROM users WHERE email = $1;
		`

	row := u.db.QueryRow(ctx, query, email)

	var i domain.User

	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.MSISDN,
		&i.EmailVerified,
		&i.MsisdnVerified,
		&i.Active,
		&i.AccountTypeId,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (u *UserRepository) GetByMSISDN(ctx context.Context, msisdn string) (*domain.User, error) {
	query := `
		SELECT * FROM users WHERE msisdn = $1;
		`

	row := u.db.QueryRow(ctx, query, msisdn)

	var i domain.User

	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.MSISDN,
		&i.EmailVerified,
		&i.MsisdnVerified,
		&i.Active,
		&i.AccountTypeId,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}

func (u *UserRepository) GetByEmailorMSISDN(ctx context.Context, params domain.GetByEmailorMSISDNParams) (*domain.User, error) {
	query := `
		SELECT * FROM users WHERE email = $1 or msisdn = $2;
		`

	row := u.db.QueryRow(ctx, query, params.Email, params.MSIDN)

	var i domain.User

	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.MSISDN,
		&i.EmailVerified,
		&i.MsisdnVerified,
		&i.Active,
		&i.AccountTypeId,
		&i.InsertedAt,
		&i.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, err
}
