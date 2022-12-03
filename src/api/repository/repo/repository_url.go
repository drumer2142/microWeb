package repo

import (
	_ "log"

	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/api/models"
)

func (r *database.RepositoryGorm) StoreWebSite(*models.Website) error {

}

func (r *database.RepositoryGorm) RetriveAllWebSites() ([]models.Website, error) {
	var err error
	allsites := []models.Website{}
	err = r.db.Find(&allsites).Error
	if err != nil {
		return nil, err
	}
	return allsites, nil
}

func (r *database.RepositoryGorm) RetriveByDomainName(domain string) ([]models.Website, error) {
	var err error
	var sitebydomain []models.Website
	err = r.db.Debug().Where("domain = ?", domain).Find(&sitebydomain).Error
	if err != nil {
		return nil, err
	}
	return sitebydomain, nil
}
