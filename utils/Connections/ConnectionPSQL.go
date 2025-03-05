package connections

import (
	"fmt"

	"github.com/jmoiron/sqlx"
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

	connPSQL := ConnectionPSQL{
		Host:     "localhost",
		Port:     "5432",
		SSLMode:  "disable",
		Username: "postgres",
		Password: "admin",
		Database: "postgres",
	}

	ConnectionMapPSQL["DEVELOPMENT"] = connPSQL
}

func (c *ConnectionPSQL) ConnectionStringPSQL() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s",
		c.Username, c.Database, c.SSLMode, c.Password, c.Host, c.Port)
}

func (c *ConnectionPSQL) ConnectionOpenPSQL() (*sqlx.DB, error) {
	fmt.Println("In connectionOpen")
	psqlDb, err := sqlx.Connect("postgres", c.ConnectionStringPSQL())
	fmt.Println("psqlDb: ", psqlDb)

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
