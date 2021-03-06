package repo

import (
	_"log"	
	"github.com/drumer2142/microWeb/src/api/models"
	"github.com/jinzhu/gorm"
  )

  type RepositoryGorm struct{
	  db *gorm.DB
  }

  func NewUrlRepo(db *gorm.DB) *RepositoryGorm {
	return &RepositoryGorm{db}
  }

  func (r *RepositoryGorm) FindAll() ([]models.Site , error){
    var err error
    allsites := []models.Site{}
    err = r.db.Find(&allsites).Error
    if err != nil{
      return nil, err
    }
    return allsites, nil
  }

  func (r *RepositoryGorm) FindByDomain(domain string) ([]models.Site, error){
    var err error
    var sitebydomain []models.Site
    err = r.db.Debug().Where("domain = ?", domain).Find(&sitebydomain).Error
    if err != nil{
      return nil, err
    }
    return sitebydomain, nil
  }
