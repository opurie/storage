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

func (s *ItemsService) GetItemById(id int) (*model.Item, error){
	var item model.Item
	if err := s.db.Preload("Category").
		Preload("Location").
		Preload("User").First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ItemsService) GetAllItems() ([]model.Item, error) {
	var items []model.Item
	if err := s.db.Preload("Category").
		Preload("Location").
		Preload("User").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ItemsService) CreateItem(item *model.Item) error {
	if err := s.db.Save(item).Error; err != nil {
		return err
	}
	
	return nil
}

func (s *ItemsService) UpdateItem(item *model.Item) (*model.Item, error) {
	if err := s.db.Model(&model.Item{}).Where("id = ?", item.ID).Updates(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemsService) DeleteItem(id int) error {
	if err := s.db.Delete(&model.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}