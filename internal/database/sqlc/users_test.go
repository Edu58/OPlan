package db_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/database"
	db "github.com/Edu58/Oplan/internal/database/sqlc"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var (
	queries *db.Queries
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx, "postgres:latest",
		postgres.WithDatabase("oplan_dev"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithSQLDriver("pgx"),
		postgres.BasicWaitStrategies(),
	)

	if err != nil {
		log.Fatalf("Error starting postgres container: %v", err)
	}

	log.Printf("Created postgres container: %s", pgContainer.GetContainerID())

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")

	if err != nil {
		log.Fatalf("Error configuring postgres container DSN: %v", err)
	}

	config := config.Config{
		DSN_URL:        dsn,
		MIGRATIONS_URL: "file://../migrations",
		LOGGER_LEVEL:   "info",
	}

	log.Printf("Configured postgres container DSN: %s", dsn)

	logger := logger.NewLoggerWithLevel(config.LOGGER_LEVEL, os.Stdout)

	pgxPool, err := database.InitDB(ctx, &config, logger)

	if err != nil {
		log.Fatalf("Error creating pgx pool: %v", err)
	}

	queries = db.New(pgxPool)

	code := m.Run()

	pgxPool.Close()
	pgContainer.Terminate(ctx)

	os.Exit(code)
}

func TestAccountTypes(t *testing.T) {
	tests := []struct {
		name   string
		setup  func(ctx context.Context) (context.Context, db.AccountType, error)
		verify func(t *testing.T, ctx context.Context, result any, err error)
	}{
		{
			name: "Create Account Type",
			setup: func(ctx context.Context) (context.Context, db.AccountType, error) {
				createAccountTypeParams := db.CreateAccountTypeParams{"Test", pgtype.Bool{Bool: true, Valid: true}}
				accType, err := queries.CreateAccountType(context.Background(), createAccountTypeParams)
				return ctx, accType, err
			},
			verify: func(t *testing.T, ctx context.Context, result any, err error) {
				require.NoError(t, err)
				r := result.(db.AccountType)
				require.NotNil(t, r.ID.Bytes)
				require.NotNil(t, r.InsertedAt)
				require.NotNil(t, r.UpdatedAt)
				require.Equal(t, r.Name, "Test")
				require.Equal(t, r.Active.Bool, true)
			},
		},
		{
			name: "Handle Duplicate Account Types",
			setup: func(ctx context.Context) (context.Context, db.AccountType, error) {
				createAccountTypeParams := db.CreateAccountTypeParams{"Test", pgtype.Bool{Bool: true, Valid: true}}
				accType, err := queries.CreateAccountType(context.Background(), createAccountTypeParams)
				return ctx, accType, err
			},
			verify: func(t *testing.T, ctx context.Context, result any, err error) {
				require.Error(t, err)
				require.ErrorContains(t, err, "duplicate key value violates unique constraint \"account_types_name_idx\"")
			},
		},
		{
			name: "GET Account Type by ID",
			setup: func(ctx context.Context) (context.Context, db.AccountType, error) {
				createAccountTypeParams := db.CreateAccountTypeParams{"Test num 2", pgtype.Bool{Bool: true, Valid: true}}
				accType, err := queries.CreateAccountType(context.Background(), createAccountTypeParams)
				ctx = context.WithValue(ctx, "ACC_TYPE_ID", accType.ID)
				return ctx, accType, err
			},
			verify: func(t *testing.T, ctx context.Context, result any, err error) {
				require.NoError(t, err)
				r := result.(db.AccountType)

				accTypeID := ctx.Value("ACC_TYPE_ID").(pgtype.UUID)
				accType, err := queries.GetAccountTypeById(context.Background(), accTypeID)

				require.NoError(t, err)
				require.Equal(t, r.ID, accType.ID)
				require.Equal(t, r.InsertedAt, accType.InsertedAt)
				require.Equal(t, r.UpdatedAt, accType.UpdatedAt)
				require.Equal(t, r.Name, accType.Name)
				require.Equal(t, r.Active, accType.Active)
			},
		},
		{
			name: "GET Account Type by Name",
			setup: func(ctx context.Context) (context.Context, db.AccountType, error) {
				return ctx, db.AccountType{}, nil
			},
			verify: func(t *testing.T, ctx context.Context, result any, err error) {
				accType, err := queries.GetAccountTypeByName(context.Background(), "Test")

				require.NoError(t, err)
				require.Equal(t, accType.Name, "Test")
			},
		},
		{
			name: "Update Account Type",
			setup: func(ctx context.Context) (context.Context, db.AccountType, error) {
				accType, err := queries.GetAccountTypeByName(context.Background(), "Test")
				return ctx, accType, err
			},
			verify: func(t *testing.T, ctx context.Context, result any, err error) {
				r := result.(db.AccountType)
				updateAccountTypeParams := db.UpdateAccountTypeByIDParams{r.ID, "Updated Test Name", r.Active}
				updatedAccType, err := queries.UpdateAccountTypeByID(context.Background(), updateAccountTypeParams)

				require.NoError(t, err)
				require.Equal(t, updatedAccType.ID, updateAccountTypeParams.ID)
				require.Equal(t, updatedAccType.Name, updateAccountTypeParams.Name)
			},
		},
		{
			name: "List Account Types",
			setup: func(ctx context.Context) (context.Context, db.AccountType, error) {
				accTypes, err := queries.ListAccountTypes(context.Background())
				ctx = context.WithValue(ctx, "ACC_TYPES", accTypes)
				return ctx, db.AccountType{}, err
			},
			verify: func(t *testing.T, ctx context.Context, result any, err error) {
				accTypes := ctx.Value("ACC_TYPES").([]db.AccountType)

				require.NoError(t, err)
				require.Len(t, accTypes, 2)
			},
		},
		{
			name: "Delete Account Type",
			setup: func(ctx context.Context) (context.Context, db.AccountType, error) {
				accType, err := queries.GetAccountTypeByName(context.Background(), "Updated Test Name")
				return ctx, accType, err
			},
			verify: func(t *testing.T, ctx context.Context, result any, err error) {
				r := result.(db.AccountType)
				deletedAccType, err := queries.DeleteAccountType(context.Background(), r.ID)

				require.NoError(t, err)
				require.Equal(t, deletedAccType.ID, r.ID)
				require.Equal(t, deletedAccType.Name, r.Name)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			ctx, accType, err := test.setup(ctx)
			test.verify(t, ctx, accType, err)
		})
	}
}
