package deployments

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Migration(conn *pgx.Conn) {
	schemaSQL, err := os.ReadFile("../../schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	tables := [2]string{"associated", "qualquer"}
	for i := range 2 {
		sqlcommand := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;", tables[i])

		if _, err = conn.Exec(context.Background(), sqlcommand); err != nil {
			log.Fatal("Erro ao aplicar schema.sql", err)
		}
	}

	_, err = conn.Exec(context.Background(), string(schemaSQL))
	if err != nil {
		log.Fatal("Erro ao aplicar schema.sql", err)
	}
}
