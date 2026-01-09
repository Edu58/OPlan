package seeds

import (
	"context"
	"log"

	"github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/brianvoe/gofakeit/v7"
)

func seedEvents(ctx context.Context, queries *sqlc.Queries) error {
	maxAge := int32(gofakeit.Int8())
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
			Capacity:         int32(gofakeit.Int8()),
			PoliciesAndRules: &policies,
			MinAge:           int32(gofakeit.Int8()),
			MaxAge:           &maxAge,
			Public:           gofakeit.Bool(),
			EventTypeID:      eventType.ID,
		},
		{
			Name:             "Digital Marketing Masterclass",
			Description:      &description,
			FromTime:         gofakeit.FutureDate(),
			ToTime:           gofakeit.FutureDate(),
			Capacity:         int32(gofakeit.Int8()),
			PoliciesAndRules: &policies,
			Public:           gofakeit.Bool(),
			MinAge:           int32(gofakeit.Int8()),
			MaxAge:           &maxAge,
			EventTypeID:      eventType.ID,
		},
		{
			Name:             "Startup Founders Meetup",
			Description:      &description,
			FromTime:         gofakeit.FutureDate(),
			ToTime:           gofakeit.FutureDate(),
			Capacity:         int32(gofakeit.Int8()),
			PoliciesAndRules: &policies,
			Public:           gofakeit.Bool(),
			MinAge:           int32(gofakeit.Int8()),
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
