package dal

import (
	"github.com/jinzhu/gorm"
	"github.com/yasar-arafat/go-project-starter/model"
)

type UserStore struct{}

func (us UserStore) GetAll() ([]model.User, error) {
	var m []model.User
	if err := DB.Find(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return m, nil
		}
		return m, err
	}
	return m, nil
}

func (us UserStore) Create(u *model.User) (err error) {
	return DB.Create(u).Error
}
