package seeds

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
)

func seedUsers(ctx context.Context, queries *sqlc.Queries) error {
	lastName := gofakeit.Name()
	active := gofakeit.Bool()

	users := []sqlc.CreateUserParams{
		{
			Email:     gofakeit.Email(),
			FirstName: gofakeit.Name(),
			LastName:  &lastName,
			Password:  "password123",
			Active:    &active,
		},
		{
			Email:     gofakeit.Email(),
			FirstName: gofakeit.Name(),
			LastName:  &lastName,
			Password:  "password1234",
			Active:    &active,
		},
		{
			Email:     gofakeit.Email(),
			FirstName: gofakeit.Name(),
			LastName:  &lastName,
			Password:  "password1235",
			Active:    &active,
		},
	}

	for _, u := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)

		_, err = queries.CreateUser(ctx, u)

		if err != nil {
			return err
		}
	}

	return nil
}
