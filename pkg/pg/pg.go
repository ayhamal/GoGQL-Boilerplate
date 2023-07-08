package pg

import (
	"context"
	"github.com/ayhamal/gogql-boilerplate/env"
	"github.com/ayhamal/gogql-boilerplate/pkg/entities"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Client instance
var pgClient *PgClient

// PostgresClient struct
type PgClient struct {
	Db *gorm.DB
}

// The GetInstance function returns a singleton instance of a Postgres client, creating a new instance
// if one does not already exist.
func GetInstance() (*PgClient, error) {
	// Check if postgresClient is not nil
	if pgClient != nil {
		return pgClient, nil
	}
	// Get environment variables
	env, err := env.New()
	// Handle error getting environment variables.
	if err != nil {
		return nil, err
	}
	// Create new postgres client instance
	pgClient, err := New(context.Background(), env)
	// Handle error creating new postgres client instance.
	if err != nil {
		return nil, err
	}
	// Return postgres client instance
	return pgClient, nil
}

// The New function initializes a new PostgreSQL client and returns it along with a nil error if
// successful.
func New(ctx context.Context, env *env.Env) (*PgClient, error) {
	// Check if postgresClient is not nil
	if pgClient != nil {
		return pgClient, nil
	}
	// Build connection string
	dsn := BuildConnectionString(env)
	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic("FATAL: Failed to connect to postgres database...")
	}
	// Reference mongo client value
	pgClient = &PgClient{Db: db}
	// Log connection success
	log.Println("[INFO] Connected to postgres database...")
	// Config database
	configDatabase(ctx, pgClient)
	// Return client with nil error
	return pgClient, nil
}

// The `CloseConnection` method is a function defined on the `PgClient` struct. It is used to close the
// connection to the PostgreSQL database.
func (pgclient *PgClient) CloseConnection() error {
	// Get database conection instance
	sqlDB, err := pgclient.Db.DB()
	// Handle error getting database connection instance.
	if err != nil {
		return err
	}
	// Check database connection instnace not nil.
	if sqlDB == nil {
		log.Println("Database connection instance is nil.")
		return nil
	}
	// Close database connection.
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}

// Method to create database collections & time series
func configDatabase(ctx context.Context, client *PgClient) {
	// Log database configuration
	log.Println("[INFO] Configuring database...")
	// Log migration users table
	log.Println("[INFO] Migrating users table...")
	// Migrate database schemas
	client.Db.AutoMigrate(&entities.User{})
	// Log migration roles table
	log.Println("[INFO] Migrating roles table...")
	// Migrate database schemas
	client.Db.AutoMigrate(&entities.Role{})
}

// Drop database tables
func (pgClient *PgClient) DropDatabaseTables() {
	// Log database configuration
	log.Println("[INFO] Dropping database tables...")
	// Rollback database schemas
	pgClient.Db.Migrator().DropTable(&entities.User{})
}

// Build connection string
func BuildConnectionString(env *env.Env) string {
	// Build connection string
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		env.PostgresDb.Host,
		env.PostgresDb.Username,
		env.PostgresDb.Password,
		env.PostgresDb.Database,
		env.PostgresDb.Port,
		env.PostgresDb.SslMode,
		env.PostgresDb.TimeZone,
	)
}
