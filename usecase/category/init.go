package category

import (
	"github.com/richardtkpd/learn-go/entity"
)

type CategoryRepo interface {
	FindById(id string) *entity.Category
}

type Usecase struct {
	categoryRepo CategoryRepo
}

//go:generate mockgen -package=category -source=./init.go -destination=_mock/init_mock.go
func New(categoryRepo CategoryRepo) *Usecase {
	return &Usecase{
		categoryRepo: categoryRepo,
	}
}
