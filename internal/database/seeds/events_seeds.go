package seeds

import (
	"context"
	"log"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/brianvoe/gofakeit/v7"
)

func seedEvents(ctx context.Context, queries *sqlc.Queries) error {
	public := false
	maxAge := gofakeit.Int32()
	description := gofakeit.ProductDescription()
	policies := `Event Rules:
	• Professional attire required
	• No outside food or drinks
	• ID required for entry`

	eventType, err := queries.GetEventTypeByName(ctx, "Technology")

	if err != nil {
		log.Printf("Error getting events seed event type: %v", err)
	}

	types := []sqlc.CreateEventParams{
		{
			Name:             "Tech Conference 2025",
			Description:      &description,
			FromTime:         gofakeit.FutureDate(),
			ToTime:           gofakeit.FutureDate(),
			Capacity:         gofakeit.Int32(),
			PoliciesAndRules: &policies,
			MinAge:           gofakeit.Int32(),
			MaxAge:           &maxAge,
			EventTypeID:      eventType.ID,
		},
		{
			Name:             "Digital Marketing Masterclass",
			Description:      &description,
			FromTime:         gofakeit.FutureDate(),
			ToTime:           gofakeit.FutureDate(),
			Capacity:         gofakeit.Int32(),
			PoliciesAndRules: &policies,
			Public:           &public,
			MinAge:           gofakeit.Int32(),
			MaxAge:           &maxAge,
			EventTypeID:      eventType.ID,
		},
		{
			Name:             "Startup Founders Meetup",
			Description:      &description,
			FromTime:         gofakeit.FutureDate(),
			ToTime:           gofakeit.FutureDate(),
			Capacity:         gofakeit.Int32(),
			PoliciesAndRules: &policies,
			MinAge:           gofakeit.Int32(),
			MaxAge:           &maxAge,
			EventTypeID:      eventType.ID,
		},
	}

	for _, t := range types {

		_, err := queries.CreateEvent(ctx, t)

		if err != nil {
			log.Printf("Error seeding event: %v", err)
			return err
		}
	}

	return nil
}
