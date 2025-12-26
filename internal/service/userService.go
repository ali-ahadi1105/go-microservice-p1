package service

import (
	"go-microservice-project/internal/domain"
	"go-microservice-project/internal/dto"
	"log"
)

type UserService struct {
}

func (s *UserService) Register(input dto.RegisterUser) (string, error) {

	log.Println("Registering user", input)

	return "my-token", nil
}

func (s *UserService) Login(input any) (string, error) {
	return "", nil
}

func (s *UserService) GetUserByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (s *UserService) GetVerificationCode(d domain.User) (int, error) {
	return 0, nil
}

func (s *UserService) VerifiedCode(id uint, code int) error {
	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s *UserService) GetProfile(id uint) (domain.User, error) {
	return domain.User{}, nil
}

func (s *UserService) UpdateProfile(id uint, input any) (domain.User, error) {
	return domain.User{}, nil
}

func (s *UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s *UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) CteateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s *UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) GetOrderById(id, uId uint) (interface{}, error) {
	return nil, nil
}
