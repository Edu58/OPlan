package seeds

import (
	"context"
	"fmt"

	"github.com/Edu58/Oplan/internal/database/sqlc"
)

func Run(ctx context.Context, queries *sqlc.Queries) error {
	if err := seedUsers(ctx, queries); err != nil {
		return fmt.Errorf("seed users: %w", err)
	}

	if err := seedEventTypes(ctx, queries); err != nil {
		return fmt.Errorf("seed event types: %w", err)
	}

	return nil
}
