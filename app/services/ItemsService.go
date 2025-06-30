package services

import (
	"main/app/models"

	"github.com/jinzhu/gorm"
)

type ItemsService struct {
	db *gorm.DB
}
func NewItemsService(db *gorm.DB) *ItemsService {
	return &ItemsService{db: db}
}

func (s *ItemsService) GetItemById(id string) (*model.Item, error){
	var item model.Item
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ItemsService) GetAllItems() ([]model.Item, error) {
	var items []model.Item
	if err := s.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemsService) CreateOrUpdateItem(item *model.Item) error {
	if err := s.db.Save(item).Error; err != nil {
		return err
	}
	return nil
}

func (s *ItemsService) DeleteItem(id string) error {
	if err := s.db.Delete(&model.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}