package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var dbURL string = fmt.Sprintf(
	"postgresql://%s:%s@%s:%s/%s",
	os.Getenv("PG_USERNAME"),
	os.Getenv("PG_PASSWORD"),
	os.Getenv("PG_HOST"),
	os.Getenv("PG_PORT"),
	os.Getenv("PG_DBNAME"),
)

func PgxConnect() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	}
	return conn
}
