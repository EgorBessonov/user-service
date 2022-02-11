//Package service represents user gRPC server handlers
package service

import (
	"context"
	"fmt"
	"github.com/EgorBessonov/user-service/internal/model"
	"github.com/EgorBessonov/user-service/internal/repository"
	userService2 "github.com/EgorBessonov/user-service/protocol"
)

//Service struct
type Service struct {
	Repository *repository.PostgresRepository
	userService2.UnimplementedUserServer
}

//NewService returns new service instance
func NewService(rps *repository.PostgresRepository) *Service {
	return &Service{Repository: rps}
}

//Registration method create new user instance in database
func (service *Service) Registration(ctx context.Context, request *userService2.RegistrationRequest) (*userService2.RegistrationResponse, error) {
	user := model.User{
		Email:    request.UserEmail,
		Password: request.UserPassword,
		Name:     request.UserName,
	}
	err := service.Repository.Save(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &userService2.RegistrationResponse{Result: fmt.Sprint("successfully added")}, nil
}

//Authentication returns user information from database
func (service *Service) Authentication(ctx context.Context, request *userService2.AuthenticationRequest) (*userService2.AuthenticationResponse, error) {
	user, err := service.Repository.Get(ctx, request.UserEmail)
	if err != nil {
		return nil, err
	}
	if user.Password != request.UserPassword {
		return nil, fmt.Errorf("invalid password")
	}
	return &userService2.AuthenticationResponse{
		UserId:   user.ID,
		UserName: user.Name,
	}, nil
}
