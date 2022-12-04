package database

import (
	"github.com/drumer2142/microWeb/src/api/models"
	"github.com/drumer2142/microWeb/src/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Storage interface {
	StoreWebSite(*models.Website) error
	RetriveAllWebSites() ([]*models.Website, error)
	RetriveByDomainName(string) (*models.Website, error)
}

type RepositoryGorm struct {
	db *gorm.DB
}

func Connect() (*RepositoryGorm, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		return nil, err
	}

	return &RepositoryGorm{
		db: db,
	}, nil

}

func (r *RepositoryGorm) StoreWebSite(*models.Website) error {

}

func (r *RepositoryGorm) RetriveAllWebSites() ([]*models.Website, error) {
	// var err error
	// allsites := []models.Website{}
	// err = r.db.Find(&allsites).Error
	// if err != nil {
	// 	return nil, err
	// }
	// return allsites, nil
}

func (r *RepositoryGorm) RetriveByDomainName(domain string) (*models.Website, error) {
	// var err error
	// var sitebydomain []models.Website
	// err = r.db.Debug().Where("domain = ?", domain).Find(&sitebydomain).Error
	// if err != nil {
	// 	return nil, err
	// }
	// return sitebydomain, nil
}
