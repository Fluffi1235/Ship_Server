package postgreSQL

import (
	"context"
	ssov1 "diplomServer/gen"
	"diplomServer/internal/models"
	"fmt"
)

func (s *Storage) GetUser(ctx context.Context, request *ssov1.GetUserRequest) (*ssov1.GetUserResponse, error) {
	const op = "repository.postgreSQL.GetRole"

	var user models.User

	query := "Select username, passwordhash, name, surname, department, rank, email, photourl  from users where username = $1"

	err := s.db.GetContext(ctx, &user, query, request.UserName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if user.UserName == "" {
		return nil, nil
	}

	req := &ssov1.GetUserResponse{
		Id:         user.Id,
		UserName:   user.UserName,
		Password:   []byte(user.PasswordHash),
		Name:       user.Name,
		Surname:    user.Surname,
		Department: user.Department,
		Rank:       user.Rank,
		Email:      user.Email,
		PhotoUrl:   user.PhotoUrl,
	}

	return req, nil
}

func (s *Storage) DeleteUser(ctx context.Context, user_id int64) (bool, error) {
	const op = "repository.postgreSQL.DeleteUser"

	query := "Delete from users where id = $1"

	_, err := s.db.ExecContext(ctx, query, user_id)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, nil
}

func (s *Storage) ChangeProfilePhoto(ctx context.Context, userName, urlPhoto string) error {
	const op = "repository.postgreSQL.ChangeProfilePhoto"
	query := "UPDATE users set photourl = $1 where username = $2"

	_, err := s.db.ExecContext(ctx, query, urlPhoto, userName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetUsersData(ctx context.Context, req *ssov1.UserName) (*ssov1.GetUsersDataResponse, error) {
	const op = "repository.postgreSQL.GetUsersData"

	var users []*ssov1.UserInfo

	query := "SELECT name, surname, department, rank, email, photourl FROM users order by id"

	err := s.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &ssov1.GetUsersDataResponse{Users: users}, nil
}

func (s *Storage) ChangePassword(ctx context.Context, req *ssov1.NewPassword) error {
	const op = "repository.postgreSQL.ChangePassword"
	query := "UPDATE users set passwordhash = $1 where username = $2"

	_, err := s.db.ExecContext(ctx, query, req.Pas, req.UserName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) ChangeEmail(ctx context.Context, req *ssov1.Email) error {
	const op = "repository.postgreSQL.ChangePassword"
	query := "UPDATE users set email = $1 where username = $2"

	_, err := s.db.ExecContext(ctx, query, req.Email, req.UserName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
