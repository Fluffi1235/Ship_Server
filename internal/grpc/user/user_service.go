package user

import (
	"bytes"
	"context"
	ssov1 "diplomServer/gen"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
	"log"
	"time"
)

var (
	accessKey  = "YCAJE0q5obszVjXWanbgi-wbR"
	secret     = "YCN4E07uyKFDGI3hCBbTXmqya__V8ttUlRroQ-Op"
	bucketName = "users-avatar"
	region     = "ru-central1"
	endpoint   = "https://storage.yandexcloud.net"
)

type UserImpl struct {
	log      *zap.SugaredLogger
	userData UserData
	tokenTTl time.Duration
}

type User interface {
	ChangeProfilePhoto(ctx context.Context, req *ssov1.ChangeProfilePhotoRequest) (*ssov1.ChangeProfilePhotoResponse, error)
	GetUserData(ctx context.Context, user *ssov1.GetUserRequest) (string, error)
	DeleteUser(ctx context.Context, req *ssov1.UserDeleteRequest) (bool, error)
	GetUsersData(ctx context.Context, req *ssov1.UserName) (*ssov1.GetUsersDataResponse, error)
	ChangePassword(ctx context.Context, req *ssov1.NewPassword) error
	ChangeEmail(ctx context.Context, req *ssov1.Email) error
}

type UserData interface {
	ChangeProfilePhoto(ctx context.Context, userName, urlPhoto string) error
	GetUser(ctx context.Context, user *ssov1.GetUserRequest) (*ssov1.GetUserResponse, error)
	DeleteUser(ctx context.Context, user_id int64) (bool, error)
	GetUsersData(ctx context.Context, req *ssov1.UserName) (*ssov1.GetUsersDataResponse, error)
	ChangePassword(ctx context.Context, req *ssov1.NewPassword) error
	ChangeEmail(ctx context.Context, req *ssov1.Email) error
}

func New(log *zap.SugaredLogger, userSaver UserData, tokenTTL time.Duration) *UserImpl {
	return &UserImpl{
		log:      log,
		userData: userSaver,
		tokenTTl: tokenTTL,
	}
}

func (u *UserImpl) GetUserData(ctx context.Context, user *ssov1.GetUserRequest) (*ssov1.GetUserResponse, error) {
	const op = "auth.GetUserRole"

	u.log.Info("get user data")

	userData, err := u.userData.GetUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	u.log.Info("get the user data: ", user.UserName)

	return userData, nil
}

func (a *UserImpl) DeleteUser(ctx context.Context, req *ssov1.UserDeleteRequest) (bool, error) {
	const op = "auth.DeleteUser"

	a.log.Info("deleting user")

	status, err := a.userData.DeleteUser(ctx, req.UserId)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("user deleted")

	return status, nil
}

func (a *UserImpl) ChangeProfilePhoto(ctx context.Context, req *ssov1.ChangeProfilePhotoRequest) (*ssov1.ChangeProfilePhotoResponse, error) {
	err := uploadFile(req.UserName, req.PhotoName, req.Photo)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://users-avatar.storage.yandexcloud.net/%s/%s", req.UserName, req.PhotoName)

	err = a.userData.ChangeProfilePhoto(ctx, req.UserName, url)
	if err != nil {

	}

	return &ssov1.ChangeProfilePhotoResponse{UrlPhoto: url}, nil
}

func (a *UserImpl) GetUsersData(ctx context.Context, req *ssov1.UserName) (*ssov1.GetUsersDataResponse, error) {
	users, err := a.userData.GetUsersData(ctx, req)
	if err != nil {

	}

	return users, nil
}

func (a *UserImpl) ChangePassword(ctx context.Context, req *ssov1.NewPassword) error {
	err := a.userData.ChangePassword(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserImpl) ChangeEmail(ctx context.Context, req *ssov1.Email) error {
	err := a.userData.ChangeEmail(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func uploadFile(userName, photoName, photoBase64 string) error {
	// Декодирование base64 строки в байты
	photoData, err := base64.StdEncoding.DecodeString(photoBase64)
	if err != nil {
		log.Println("Failed to decode base64 string")
		return err
	}

	// Создаем новую сессию AWS с пользовательскими учетными данными
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secret, ""),
		Endpoint:    aws.String(endpoint),
	})
	if err != nil {
		log.Println("Failed to create AWS session: ", err)
		return err
	}
	svc := s3.New(sess)

	// Формируем ключ объекта S3
	key := userName + "/" + photoName

	// Загружаем файл в Amazon S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(photoData),
		ContentType: aws.String("image/jpeg"), // Установка MIME типа, если известен
	})
	if err != nil {
		log.Println("Failed to upload file to S3: ", err)
		return err
	}

	log.Println("File uploaded successfully")
	return nil
}
