package service

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (s *UserServiceImpl) UpdateUser(input web.FormUpdateUserInput) (domain.User, error) {
	user, err := s.UserRepository.FindByID(input.ID)
	helper.PanicIfError(err)

	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email

	update, err := s.UserRepository.Update(user)
	helper.PanicIfError(err)

	return update, nil
}

func (s *UserServiceImpl) FindAllUsers() ([]domain.User, error) {
	users, err := s.UserRepository.FindAll()
	helper.PanicIfError(err)

	return users, nil
}

func (s *UserServiceImpl) FindUserByID(id int) (domain.User, error) {
	findByID, err := s.UserRepository.FindByID(id)
	helper.PanicIfError(err)

	if findByID.ID == 0 {
		return findByID, errors.New("User not found")
	}

	return findByID, nil
}

func (s *UserServiceImpl) UploadAvatar(userID int, fileLocation string) (domain.User, error) {
	findByID, err := s.UserRepository.FindByID(userID)
	helper.PanicIfError(err)

	findByID.Avatar = fileLocation

	update, err := s.UserRepository.Update(findByID)
	helper.PanicIfError(err)

	return update, nil
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

	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	if findByEmail.Email == input.Email {
		return user, errors.New("Email not available")
	}

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
