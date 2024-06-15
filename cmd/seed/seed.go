package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chaninlaw/auth/pkg"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pkg.LoadEnv(".env")
	
	conn, err := pkg.DBConnectPool()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return
	}

	err = SeedUser(conn)
	if err != nil {
		log.Fatalf("Failed to seed database: %v\n", err)
	}

	fmt.Println("Database seeded successfully")
	os.Exit(0)
}

func SeedUser(pool *pgxpool.Pool) error {
	ctx := context.Background()
	dbpool, err := pool.Acquire(ctx)
	if err != nil {
		fmt.Println("Failed to start transaction")
		return err
	}
	defer dbpool.Release()

	// Create tables
	_, err = dbpool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			is_admin BOOLEAN DEFAULT false
		);
	`)
	if err != nil {
		fmt.Println("Failed to create users table")
		return err
	}

	// Hash password for the admin user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to hash password")
		return err
	}

	// Insert admin user
	_, err = dbpool.Exec(ctx, `
		INSERT INTO users (username, password, is_admin) VALUES ($1, $2, $3)
		ON CONFLICT (username) DO NOTHING;`, "admin", string(hashedPassword), true)
	if err != nil {
		fmt.Println("Failed to insert admin user")
		return err
	}

	return nil
}