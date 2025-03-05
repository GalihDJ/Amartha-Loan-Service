package connections

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ConnectionPSQL struct {
	Host     string
	Port     string
	SSLMode  string
	Username string
	Password string
	Database string
}

var ConnectionMapPSQL map[string]ConnectionPSQL

func init() {
	ConnectionMapPSQL = make(map[string]ConnectionPSQL)

	if os.Getenv("LOAN_SERVICE_PLATFORM_ENVIRONMENT") == "DEVELOPMENT" || os.Getenv("LOAN_SERVICE_PLATFORM_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	psqlHost := os.Getenv("POSTGRES_DB_HOST")
	psqlPort := os.Getenv("POSTGRES_DB_PORT")
	psqlSSLMode := os.Getenv("POSTGRES_DB_SSL_MODE")
	psqlUsername := os.Getenv("POSTGRES_DB_USERNAME")
	psqlPassword := os.Getenv("POSTGRES_DB_PASSWORD")
	psqlDatabase := os.Getenv("POSTGRES_DB_DATABASE")

	connPSQL := ConnectionPSQL{
		Host:     psqlHost,
		Port:     psqlPort,
		SSLMode:  psqlSSLMode,
		Username: psqlUsername,
		Password: psqlPassword,
		Database: psqlDatabase,
	}

	ConnectionMapPSQL[os.Getenv("LOAN_SERVICE_PLATFORM_ENVIRONMENT")] = connPSQL
}

func (c *ConnectionPSQL) ConnectionStringPSQL() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s",
		c.Username, c.Database, c.SSLMode, c.Password, c.Host, c.Port)
}

func (c *ConnectionPSQL) ConnectionOpenPSQL() (*sqlx.DB, error) {

	psqlDb, err := sqlx.Connect("postgres", c.ConnectionStringPSQL())

	if err != nil {
		return nil, err
	}

	return psqlDb, PingConnectionPSQL(psqlDb)
}

func PingConnectionPSQL(psqlDb *sqlx.DB) error {
	err := psqlDb.Ping()
	if err != nil {
		return err
	}

	return nil
}
