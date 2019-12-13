package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

// Initializes connection to pakim database
func connectDb(applicationName string) (conn *pgx.Conn) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = applicationName
	connConfig := pgx.ConnConfig{
		// TODO: store in consts or configmap
		User:              os.Getenv("PG_USER"),
		Password:          os.Getenv("PG_PASSWORD"),
		Host:              os.Getenv("PG_HOST"),
		Port:              5432,
		Database:          os.Getenv("PG_DATABASE"),
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish connection: %v\n", err)
		os.Exit(1)
	}
	return conn
}
