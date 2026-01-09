package seeds

import (
	"context"
	"errors"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
)

func seedEventTypes(ctx context.Context, queries *sqlc.Queries) error {

	types := []sqlc.CreateEventTypeParams{
		{
			Name:   "Technology",
			Active: true,
		},
		{
			Name:   "Business",
			Active: true,
		},
		{
			Name:   "Networking",
			Active: true,
		},
	}

	for _, t := range types {

		_, err := queries.CreateEventType(ctx, t)

		if err != nil {
			var pgErr *pgconn.PgError

			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					return nil
				}
			}
			return err
		}
	}

	return nil
}
