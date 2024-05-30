package auth

import (
	ssov1 "diplomServer/gen"
	"diplomServer/internal/lib/jwt"
	"diplomServer/internal/models"
	"diplomServer/internal/repository"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type AuthImpl struct {
	log      *zap.SugaredLogger
	userAuth UserAuth
	tokenTTl time.Duration
}

type Auth interface {
	RegisterNewUser(ctx context.Context, user *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error)
	Login(ctx context.Context, user *ssov1.LoginRequest) (*ssov1.LoginResponse, error)
}

type UserAuth interface {
	SaveUser(ctx context.Context, user *ssov1.RegisterRequest) (uid *ssov1.RegisterResponse, err error)
	Login(ctx context.Context, user *ssov1.LoginRequest) (*models.User, error)
}

// New returns a new instance
func New(log *zap.SugaredLogger, userSaver UserAuth, tokenTTL time.Duration) *AuthImpl {
	return &AuthImpl{
		log:      log,
		userAuth: userSaver,
		tokenTTl: tokenTTL,
	}
}

func (a *AuthImpl) RegisterNewUser(ctx context.Context, user *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	const op = "auth.RegisterNewUser"

	userName, err := a.userAuth.SaveUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			a.log.Warn("user alreadu exists", err.Error())

			return nil, fmt.Errorf("%s: %w", op, ErrUserExists)
		}

		a.log.Error("failed to save user", err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("user registered")

	return userName, nil
}

func (a *AuthImpl) Login(ctx context.Context, userData *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	const op = "auth.UserName"

	user, err := a.userAuth.Login(ctx, userData)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			a.log.Warn("user not found", err.Error())
			return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", err.Error())
		return nil, fmt.Errorf("%s: %w", op, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), userData.Password)
	if err != nil {
		a.log.Error("invalid credentials", err.Error())
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	a.log.Info("user logged in successfully")

	token, err := jwt.NewToken(user, a.tokenTTl)
	if err != nil {
		a.log.Error("failed to generate token", err.Error())

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &ssov1.LoginResponse{Token: token}, nil
}
