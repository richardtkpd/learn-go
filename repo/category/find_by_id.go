package category

import "github.com/richardtkpd/learn-go/entity"

func FindById(id string) *entity.Category {
	return &entity.Category{
		Id:   123,
		Name: "Richie Ren",
	}
}
