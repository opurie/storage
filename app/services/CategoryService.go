package services
import (
	"main/app/models"

	"github.com/jinzhu/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (s *CategoryService) GetCategory(id string) (*model.Category, error) {
	var category model.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (s *CategoryService) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := s.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) CreateOrUpdateCategory(category *model.Category) error {
	if err := s.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) DeleteCategory(id string) error {
	if err := s.db.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
