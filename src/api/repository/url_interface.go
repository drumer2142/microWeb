package repository

import (
	"github.com/drumer2142/microWeb/src/api/models"
)

type UrlRepository interface {
	FindAll() ([]models.Site, error)
	FindByDomain(string) ([]models.Site, error)
}