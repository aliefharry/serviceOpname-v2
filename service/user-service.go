package service

import (
	// "fmt"
	"log"

	"github.com/mashingan/smapping"
	"serviceOpname-v2/config/entity"
	"serviceOpname-v2/controller/dto"
	"serviceOpname-v2/repository"
)

type UserService interface{
	Update(user dto.UserUpdateDTO) entity.User
	Profile(UserID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService {
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(UserID string) entity.User {
	return service.userRepository.ProfileUser(UserID)
}