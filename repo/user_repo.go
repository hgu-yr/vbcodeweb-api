package repo

import (
	"vbCodeWeb/model"

	"gorm.io/gorm"
)

// UserRepo 管理员账号数据访问。
type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// GetByUsername 按用户名查询账号。
func (r *UserRepo) GetByUsername(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.db.Where("username = ?", username).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// Create 新建账号。
func (r *UserRepo) Create(u *model.User) error {
	return r.db.Create(u).Error
}
