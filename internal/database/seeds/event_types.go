package seeds

import (
	"context"

	"github.com/Edu58/Oplan/internal/database/sqlc"
)

func seedEventTypes(ctx context.Context, queries *sqlc.Queries) error {
	active := true

	types := []sqlc.CreateEventTypeParams{
		{
			Name:   "Technology",
			Active: &active,
		},
		{
			Name:   "Business",
			Active: &active,
		},
		{
			Name:   "Networking",
			Active: &active,
		},
	}

	for _, t := range types {

		_, err := queries.CreateEventType(ctx, t)

		if err != nil {
			return err
		}
	}

	return nil
}
