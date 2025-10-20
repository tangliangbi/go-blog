package dao

import (
	"errors"

	"com.tang.blog/repository/db"
	"com.tang.blog/repository/model"
	"com.tang.blog/utils/logger"
	"com.tang.blog/utils/pagination"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{db.GetDB()}
}

func (dao *UserDao) InsertUser(user *model.User) error {
	if dao == nil {
		logger.Log.Error("数据库连接异常")
		return errors.New("数据库连接异常")
	}

	err := dao.DB.Create(user).Error
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
			"password": user.Password,
			"error":    err.Error(),
		}).Error("Failed to insert user into database")
		return err
	}

	return nil
}

func (dao *UserDao) DeleteByUsername(username string) error {
	err := dao.DB.Where("username = ?", username).Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (dao *UserDao) UpdateUser(user *model.User) error {
	err := dao.DB.Model(user).Where("username = ?", user.Username).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDao) QueryByUsername(username string) (*model.User, error) {
	var user model.User
	err := dao.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (dao *UserDao) ListPaginationWithKeywords(pageParams *pagination.QueryParams, username string) (*pagination.PageResult, error) {
	var users []model.User
	query := dao.DB.Model(&users)
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}

	result, err := pagination.PageQuery(query, &users, pageParams)
	if err != nil {
		return nil, err
	}

	return result, nil
}
