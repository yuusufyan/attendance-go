package repositories

import (
	"attendance-go/src/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(create *models.UserModels) error {
	return r.db.Create(create).Error
}

func (r *UserRepository) GetByID(id uint) (*models.UserModels, error) {
	var user models.UserModels
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.UserModels, error) {
	var user models.UserModels
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUsername(username string) (*models.UserModels, error) {
	var user models.UserModels
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll(limit, offset int, search, sortBy, orderBy string) ([]models.UserModels, int64, error) {
	var user []models.UserModels
	var total int64

	baseQuery := r.db.Model(&models.UserModels{})
	baseQuery = baseQuery.Where("is_active = ?", true)

	if search != "" {
		baseQuery = baseQuery.Where("username LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting
	allowedSortColumns := map[string]bool{
		"id":         true,
		"name":       true,
		"created_at": true,
	}

	if !allowedSortColumns[sortBy] {
		sortBy = "id" // default kalau user masukin kolom aneh
	}

	if orderBy != "asc" && orderBy != "desc" {
		orderBy = "asc"
	}

	orderClause := fmt.Sprintf("%s %s", sortBy, orderBy)

	err := baseQuery.
		Order(orderClause).
		Limit(limit).
		Offset(offset).
		Find(&user).Error

	if err != nil {
		return nil, 0, err
	}

	return user, total, nil
}
