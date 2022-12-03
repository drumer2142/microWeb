package repository

import (
	"github.com/drumer2142/microWeb/src/api/models"
)

type UrlRepository interface {
	FindAll() ([]models.Website, error)
	FindByDomain(string) ([]models.Website, error)
}
