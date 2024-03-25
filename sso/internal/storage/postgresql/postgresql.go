package postgresql

import (
	"context"
	"github.com/17HIERARCH70/messageService/sso/internal/config"
	"github.com/17HIERARCH70/messageService/sso/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

// Connect establishes a connection to the PostgreSQL database.
func Connect(cfg *config.PostgresSQLConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewStorage creates a new Storage instance.
func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// SaveUser inserts a new user into the database.
func (p *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const query = `INSERT INTO users (email, pass_hash) VALUES ($1, $2) RETURNING id`
	var id int64
	err := p.db.QueryRowContext(ctx, query, email, passHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// User retrieves a user by their email.
func (p *Storage) User(ctx context.Context, email string) (models.User, error) {
	const query = `SELECT id, email, pass_hash FROM users WHERE email = $1`
	var user models.User
	err := p.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Storage) App(ctx context.Context, appID int) (models.App, error) {
	const query = `SELECT id, name FROM apps WHERE id = $1`
	var app models.App
	err := s.db.GetContext(ctx, &app, query, appID)
	if err != nil {
		// Handle error, could be sql.ErrNoRows or any other DB errors
		return models.App{}, err
	}
	return app, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const query = `SELECT EXISTS(SELECT 1 FROM admins WHERE user_id = $1)`
	var isAdmin bool
	err := s.db.QueryRowContext(ctx, query, userID).Scan(&isAdmin)
	if err != nil {
		return false, err
	}
	return isAdmin, nil
}
