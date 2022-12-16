package order

import (
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"google.golang.org/grpc/codes"

	"gitlab.mapcard.pro/external-map-team/api-proto/payment/api"
)

//go:embed migrations/*.sql
var fs embed.FS

// version defines the current migration version. This ensures the app
// is always compatible with the version of the database.
const version = 1

// Migrate migrates the Postgres schema to the current version.
func validateSchema(db *sql.DB, scheme string) error {
	sourceInstance, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	var driverInstance database.Driver
	switch scheme {
	case "postgres", "postgresql":
		driverInstance, err = postgres.WithInstance(db, new(postgres.Config))
	default:
		return fmt.Errorf("unknown scheme: %q", scheme)
	}
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", sourceInstance, scheme, driverInstance)
	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}

func orderPostgresToProto(pgOrder Order) (*api.Order, error) {
	bytes, err := pgOrder.OrderRequest.MarshalJSON()
	if err != nil {
		return nil, Log().StatusErrorf(codes.Internal, "Failed to marshallling order request: %v", err)
	}
	var orderRequest api.OrderRequest
	err = json.Unmarshal(bytes, &orderRequest)
	if err != nil {
		return nil, Log().StatusErrorf(codes.Internal, "Failed to unmarshallling order request: %v", err)
	}
	return &api.Order{
		Success:      true,
		OrderId:      pgOrder.OrderID,
		Key:          pgOrder.SellerID,
		Amount:       orderRequest.Amount,
		Type:         orderRequest.Type,
		Rrn:          pgOrder.Rrn,
		OrderRequest: &orderRequest,
	}, nil
}

func orderRawPostgresToProto(pgOrderRaw GetOrderRow) (*api.Order, error) {
	var orderRequest api.OrderRequest
	err := json.Unmarshal(pgOrderRaw.OrderRequest, &orderRequest)
	if err != nil {
		return nil, Log().StatusErrorf(codes.Internal, "Failed to unmarshallling order request: %v", err)
	}

	return &api.Order{
		Success:      true,
		OrderId:      pgOrderRaw.OrderID,
		Key:          pgOrderRaw.SellerID,
		Amount:       orderRequest.Amount,
		Type:         orderRequest.Type,
		Rrn:          pgOrderRaw.Rrn,
		OrderRequest: &orderRequest,
	}, nil
}

const defaultPgDriver = "postgres"
const defaultPgPort = "5432"

func CreateURL_FromEnvParts() (string, error) {
	var pgUrl string = ""
	pgHost := os.Getenv("DB_HOST")
	if pgHost == "" {
		return "", fmt.Errorf("Postgres host must be set")
	}
	pgDriver := os.Getenv("DB_DRIVER")
	if pgDriver == "" {
		pgDriver = defaultPgDriver
	}
	pgUser := os.Getenv("DB_USER")
	if pgUser == "" {
		return "", fmt.Errorf("Postgres user must be set")
	}
	pgPassword := os.Getenv("DB_PASSWORD")
	if pgPassword == "" {
		return "", fmt.Errorf("Postgres password must be set")
	}
	pgDbName := os.Getenv("DB_NAME")
	if pgDbName == "" {
		return "", fmt.Errorf("Postgres database name must be set")
	}
	pgPort := os.Getenv("DB_PORT")
	if pgPort == "" {
		pgPort = defaultPgPort
	}
	pgUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		pgUser,
		pgPassword,
		pgHost,
		pgPort,
		pgDbName,
	)

	return pgUrl, nil
}
