package services

import (
	"context"

	"gorm.io/gorm"
)

type HealthService struct {
	db *gorm.DB
}

func NewHealthService(db *gorm.DB) *HealthService {
	return &HealthService{db: db}
}

func (s *HealthService) CheckReadiness(ctx context.Context) (map[string]string, bool) {
	status := make(map[string]string)
	isHealthy := true

	sqlDB, err := s.db.DB()
	if err != nil {
		status["database"] = "down"
		isHealthy = false
	} else {
		if err := sqlDB.PingContext(ctx); err != nil {
			status["database"] = "down"
			isHealthy = false
		} else {
			status["database"] = "ready"
		}
	}

	return status, isHealthy
}
