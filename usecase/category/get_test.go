package category

import (
	"testing"

	mock "github.com/richardtkpd/learn-go/usecase/category/_mock"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/richardtkpd/learn-go/entity"
)

func TestUsecase_Get(t *testing.T) {
	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *entity.Category
		mock    func(ctrl *gomock.Controller) CategoryRepo
	}{
		{
			name: "success",
			args: args{
				id: "123",
			},
			wantErr: false,
			want: &entity.Category{
				Id:   123,
				Name: "Richard Ren",
			},
			mock: func(ctrl *gomock.Controller) CategoryRepo {
				m := mock.NewMockCategoryRepo(ctrl)
				m.EXPECT().FindById(gomock.Any()).Return(&entity.Category{
					Id:   123,
					Name: "Richard Ren",
				})
				return m
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				categoryRepo: tt.mock(gomock.NewController(t)),
			}
			got, err := uc.Get(tt.args.id)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
