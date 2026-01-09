package seeds

import (
	"context"
	"fmt"

	"github.com/Edu58/Oplan/internal/database/sqlc"
)

// Run seeds based on type
func Seed(ctx context.Context, queries *sqlc.Queries, seedType string) error {
	switch seedType {
	case "all":
		if err := seedUsers(ctx, queries); err != nil {
			return fmt.Errorf("seed users: %w", err)
		}

		if err := seedEventTypes(ctx, queries); err != nil {
			return fmt.Errorf("seed event types: %w", err)
		}

		if err := seedEvents(ctx, queries); err != nil {
			return fmt.Errorf("seed events: %w", err)
		}

		return nil
	case "users":
		return seedUsers(ctx, queries)
	case "events-types":
		return seedEventTypes(ctx, queries)
	case "events":
		return seedEvents(ctx, queries)
	default:
		return fmt.Errorf("unknown seed type: %s", seedType)
	}
}
