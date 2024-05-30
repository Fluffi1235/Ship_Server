package postgreSQL

import (
	"context"
	ssov1 "diplomServer/gen"
	"diplomServer/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repo interface {
	UserRepo
	Machine
}

type Storage struct {
	db *sqlx.DB
}

func New(cfg *models.Config) (Repo, error) {
	db, err := sqlx.Connect("postgres", cfg.ConnectDB)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

type UserRepo interface {
	SaveUser(ctx context.Context, user *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error)
	ChangeProfilePhoto(ctx context.Context, userName, urlPhoto string) error
	Login(ctx context.Context, user *ssov1.LoginRequest) (*models.User, error)
	GetUser(ctx context.Context, user *ssov1.GetUserRequest) (*ssov1.GetUserResponse, error)
	DeleteUser(ctx context.Context, user int64) (bool, error)
	GetUsersData(ctx context.Context, req *ssov1.UserName) (*ssov1.GetUsersDataResponse, error)
	ChangePassword(ctx context.Context, req *ssov1.NewPassword) error
	ChangeEmail(ctx context.Context, req *ssov1.Email) error
}
