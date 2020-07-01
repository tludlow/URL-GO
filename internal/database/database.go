package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func New(ctx context.Context) (error) {
	databaseURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "testpassword", "postgres")

	//Connect to the database
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return fmt.Errorf("could not connect to the database: %v", err.Error())
	}

	//Check that the connection is good
	err = conn.Ping(ctx)
	if err != nil {
		return fmt.Errorf("could not connect to the database, bad ping response: %v", err.Error())
	}

	DB = conn
	return nil
}