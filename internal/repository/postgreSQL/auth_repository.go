package postgreSQL

import (
	"context"
	ssov1 "diplomServer/gen"
	"diplomServer/internal/models"
	"fmt"
)

const AvatarUrl = "https://storage.yandexcloud.net/users-avatar/trash/defaultAvatar.jpg"

func (s *Storage) SaveUser(ctx context.Context, user *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	const op = "repository.postgreSQL.SaveUser"
	var userName string

	query := "insert into users(username, passwordhash, name, surname, department, rank, email, photourl) values($1, $2, $3, $4, $5, $6, $7, $8) returning username"

	err := s.db.GetContext(ctx, &userName, query, user.UserName, user.Password, user.Name, user.Surname, user.Department, user.Rank, user.Email, AvatarUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &ssov1.RegisterResponse{UserName: userName}, nil
}

func (s *Storage) Login(ctx context.Context, userData *ssov1.LoginRequest) (*models.User, error) {
	const op = "repository.postgreSQL.Login"

	var user models.User

	query := "Select id, userName, passwordHash, name, surname, department, rank, email, photourl from users where userName = $1"

	err := s.db.GetContext(ctx, &user, query, userData.UserName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
