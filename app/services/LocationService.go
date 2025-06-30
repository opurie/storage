package services

import (
	model "main/app/models"

	"github.com/jinzhu/gorm"
)

type LocationService struct {
	db *gorm.DB
}

func NewLocationService(db *gorm.DB) *LocationService {
	return &LocationService{db: db}
}

func (s *LocationService) GetLocation(id int) (*model.Location, error) {
	var location model.Location
	if err := s.db.First(&location, id).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func (s *LocationService) GetAllLocations() ([]model.Location, error) {
	var locations []model.Location
	if err := s.db.Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (s *LocationService) CreateLocation(location *model.Location) error {
	if err := s.db.Save(location).Error; err != nil {
		return err
	}
	return nil
}

func (s *LocationService) UpdateLocation(location *model.Location) error {
	if err := s.db.Model(&model.Location{}).Where("id = ?", location.ID).Updates(location).Error; err != nil {
		return err
	}
	return nil
}

func (s *LocationService) DeleteLocation(id int) error {
	if err := s.db.Delete(&model.Location{}, id).Error; err != nil {
		return err
	}
	return nil
}
