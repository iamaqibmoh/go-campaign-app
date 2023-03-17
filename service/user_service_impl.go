package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (s *UserServiceImpl) CheckEmailAvailability(input web.CheckEmailInput) (bool, error) {
	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	if findByEmail.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *UserServiceImpl) LoginUser(input web.LoginUserInput) (domain.User, error) {
	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return domain.User{}, err
	}

	if findByEmail.ID == 0 {
		if err != nil {
			return domain.User{}, errors.New("Login failed")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.PasswordHash), []byte(input.Password))
	if err != nil {
		return domain.User{}, errors.New("Login failed")
	}

	return findByEmail, nil
}

func (s *UserServiceImpl) RegisterUser(input web.RegisterUserInput) (domain.User, error) {
	//mapping input request to domain.User
	user := domain.User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	//call repository.Save
	save, err := s.UserRepository.Save(user)
	if err != nil {
		return user, err
	}

	return save, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}
