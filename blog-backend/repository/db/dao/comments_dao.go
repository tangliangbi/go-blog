package dao

import (
	"errors"

	"com.tang.blog/repository/db"
	"com.tang.blog/repository/model"
	"gorm.io/gorm"
)

type CommentDao struct {
	DB *gorm.DB
}

func NewCommentDao() *CommentDao {
	return &CommentDao{DB: db.GetDB()}
}

func (c *CommentDao) InsertComment(comment *model.Comment) error {
	err := c.DB.Create(comment).Error
	return err
}

func (dao *CommentDao) DeleteById(comment *model.Comment) error {
	result := dao.DB.Where("id = ? AND user_id = ?", comment.ID, comment.UserID).Delete(&model.Comment{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		var count int64
		dao.DB.Model(&model.Comment{}).Where("id = ?", comment.ID).Count(&count)
		if count > 0 {
			return ErrPermissionDenied
		}
		return ErrRecordNotFound
	}

	return nil
}

func (dao *CommentDao) UpdateById(comment *model.Comment) error {
	return dao.DB.Model(&model.Comment{}).Where("id = ?", comment.ID).Updates(comment).Error
}

func (dao *CommentDao) QueryById(comment *model.Comment) error {
	var existingComment model.Comment
	err := dao.DB.Where("id = ?", comment.ID).First(&existingComment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("comment not found")
		}
		return err
	}

	return nil
}

func (dao *CommentDao) QueryByPostId(postId uint) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := dao.DB.Where("post_id = ?", postId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
