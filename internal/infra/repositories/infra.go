package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var config = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

var dbURL string = fmt.Sprintf(
	"postgresql://%s:%s@%s:%s/%s",
	"postgres",
	"postgres",
	"localhost",
	"5432",
	"postgres",
)

func PgxConnect() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	}
	return conn
}
