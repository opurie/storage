package services

import (
	"main/app/models"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) CreateUser(user *model.User) error {
	if err := s.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(user *model.User) error {
	if err := s.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id uint) error {
	if err := s.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
