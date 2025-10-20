package dao

import (
	"context"
	"errors"

	"com.tang.blog/repository/db"
	"com.tang.blog/repository/model"
	"com.tang.blog/utils/logger"
	"com.tang.blog/utils/pagination"
	"gorm.io/gorm"
)

type PostsDao struct {
	DB *gorm.DB
}

func NewPostsDao(ctx context.Context) *PostsDao {
	return &PostsDao{DB: db.GetDB()}
}

// NewPostDao 无context的构造函数，用于在service层直接调用
func NewPostDao() *PostsDao {
	return &PostsDao{DB: db.GetDB()}
}

// InsertPost 创建文章
func (dao *PostsDao) InsertPost(post *model.Post) error {
	if dao.DB == nil {
		logger.Log.Error("Database connection is nil")
		return errors.New("database connection is not initialized")
	}

	logger.Log.WithFields(map[string]interface{}{
		"title":  post.Title,
		"userID": post.UserID,
	}).Debug("Inserting post into database")

	err := dao.DB.Create(post).Error
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"title":  post.Title,
			"userID": post.UserID,
			"error":  err.Error(),
		}).Error("Failed to insert post into database")
		return err
	}

	logger.Log.WithFields(map[string]interface{}{
		"postID": post.ID,
		"title":  post.Title,
		"userID": post.UserID,
	}).Info("Successfully inserted post into database")

	return nil
}

func (dao *PostsDao) GetPostById(id uint) (*model.Post, error) {
	var post model.Post
	err := dao.DB.Where("id = ?", id).Preload("Comments").Preload("Comments.User").First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		logger.Log.WithFields(map[string]interface{}{
			"postID": id,
			"error":  err.Error(),
		}).Error("Failed to query post by ID")
		return nil, err
	}
	return &post, nil
}

func (dao *PostsDao) UpdatePost(post *model.Post) error {
	err := dao.DB.Model(&model.Post{}).Where("id = ? AND user_id = ?", post.ID, post.UserID).Updates(post).Error
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"postID": post.ID,
			"userID": post.UserID,
			"error":  err.Error(),
		}).Error("更新文章失败")
		return err
	}
	return nil
}

func (dao *PostsDao) DeletePostById(post *model.Post) error {
	err := dao.DB.Delete(&model.Post{ID: post.ID, UserID: post.UserID}).Error
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"postID": post.ID,
			"title":  post.Title,
			"userID": post.UserID,
			"error":  err.Error(),
		}).Error("删除失败")

		return err
	}

	return nil
}

func (dao *PostsDao) ListPostsWithPagination(paginationParams *pagination.QueryParams, title, content string) (*pagination.PageResult, error) {
	var posts []model.Post

	query := dao.DB.Model(&posts).Preload("Comments").Preload("Comments.User")
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if content != "" {
		query = query.Where("content LIKE ?", "%"+content+"%")
	}

	result, err := pagination.PageQuery(query, &posts, paginationParams)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Error("Failed to query posts with pagination")
		return nil, err
	}

	return result, nil
}
