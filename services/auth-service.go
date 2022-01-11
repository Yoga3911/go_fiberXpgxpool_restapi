package services

import (
	"coba/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthService interface {
	CreateUser(models.Register, context.Context) error
	VerifyUser(models.Login, context.Context) error
}

type authService struct {
	db *pgxpool.Pool
}

func NewAuthService(db *pgxpool.Pool) AuthService {
	return &authService{
		db: db,
	}
}

const create = `INSERT INTO users (name,email,password,gender_id,create_at,update_at)
				VALUES ($1, $2, $3, $4, now(), now())`

func (a *authService) CreateUser(r models.Register, ctx context.Context) error {
	_, err := a.db.Exec(ctx, create, r.Name, r.Email, r.Password, r.GenderID)
	if err != nil {
		return err
	}
	return nil
}

const find = `SELECT email, password FROM users WHERE email = $1`

func (a *authService) VerifyUser(l models.Login, ctx context.Context) error {
	var login models.Login
	result := a.db.QueryRow(ctx, find, l.Email)
	err := result.Scan(&login.Email, &login.Password)
	if err != nil {
		return err
	}
	
	if l.Password != login.Password {
		return fmt.Errorf("invalid credential")
	} 
	
	return nil
}