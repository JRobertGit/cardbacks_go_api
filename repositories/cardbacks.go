package repositories

import (
	"go-api/data"
	"go-api/models"
)

type CardBackRepository interface {
	Create(cardback *models.CardBack) error
	GetAll() (*[]models.CardBack, error)
	FindByID(id int) (*models.CardBack, error)
	Exists(slug string) bool
}

type cardBackRepository struct {
	*data.FileSource
}

func NewCardBackRepository(dataSource *data.FileSource) CardBackRepository {
	return &cardBackRepository{dataSource}
}

func (cbr *cardBackRepository) Create(cardback *models.CardBack) error {
	if !cbr.Exists(cardback.Slug) {
		data := append(*cbr.Data, *cardback)
		cbr.Data = &data
		go cbr.Save()
	}

	return nil
}

func (cbr *cardBackRepository) GetAll() (*[]models.CardBack, error) {
	return cbr.Data, nil
}

func (cbr *cardBackRepository) FindByID(id int) (*models.CardBack, error) {
	for _, cardBack := range *cbr.Data {
		if cardBack.ID == id {
			return &cardBack, nil
		}
	}

	return nil, nil
}

func (cbr *cardBackRepository) Exists(slug string) bool {
	for _, cardBack := range *cbr.Data {
		if cardBack.Slug == slug {
			return true
		}
	}

	return false
}
