package repo

import (
	"github.com/drumer2142/microWeb/src/api/models"
	"gorm.io/gorm"
)

type RepositoryGorm struct {
	db *gorm.DB
}

func NewUrlRepo(db *gorm.DB) *RepositoryGorm {
	return &RepositoryGorm{db: db}
}

func (r *RepositoryGorm) FindAll() ([]models.Site, error) {
	var allsites []models.Site
	if err := r.db.Find(&allsites).Error; err != nil {
		return nil, err
	}
	return allsites, nil
}

func (r *RepositoryGorm) FindByDomain(domain string) ([]models.Site, error) {
	var sitebydomain []models.Site
	if err := r.db.Where("domain = ?", domain).Find(&sitebydomain).Error; err != nil {
		return nil, err
	}
	return sitebydomain, nil
}
