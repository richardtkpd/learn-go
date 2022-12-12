package category

import (
	"errors"

	"github.com/richardtkpd/learn-go/entity"
)

func (uc *Usecase) Get(id string) (*entity.Category, error) {
	category := uc.categoryRepo.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
