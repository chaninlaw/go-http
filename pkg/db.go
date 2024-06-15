package pkg

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			return nil, err
	}
	return conn, nil
}

func DBConnectPool() (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			return nil, err
	}

	 poolConfig.MaxConnLifetime = time.Hour
	 poolConfig.MaxConnIdleTime = 15 * time.Minute
	 poolConfig.MinConns = 5
	 poolConfig.MaxConns = 20

	 pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	 if err != nil {
			 fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			 return nil, err
	 }

	 return pool, nil
}