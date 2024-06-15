package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"` // omit password from JSON output
	IsAdmin  bool   `json:"is_admin"`
}

type UserRepository interface {
	  GetAll(ctx context.Context) ([]User, error)
    GetByID(ctx context.Context, id string) (*User, error)
    Create(ctx context.Context, user *User) error
}

type userRepository struct {
    db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) GetAll(ctx context.Context) ([]User, error) {
    var users []User
    rows, err := r.db.Query(ctx, "SELECT id, username, is_admin FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.Username, &u.IsAdmin)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*User, error) {
    var user User
    err := r.db.QueryRow(ctx, "SELECT id, username, is_admin FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.IsAdmin)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *User) error {
    _, err := r.db.Exec(
			ctx, 
			"INSERT INTO users (username, password, is_admin) VALUES ($1, $2, $3)", 
			user.Username, user.Password, user.IsAdmin,
		)
    return err
}
