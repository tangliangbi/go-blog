package pagination

import (
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
	Pages    int   `json:"pages"`
}

type QueryParams struct {
	Page     int `form:"page,default=1"`
	PageSize int `form:"page_size,default=10"`
}

type PageResult struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

func NewPagination(page, pageSize int, total int64) *Pagination {
	pages := 0
	if pageSize > 0 {
		pages = int(math.Ceil(float64(total) / float64(pageSize)))
	}

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Pages:    pages,
	}
}

func (p *QueryParams) Offset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.PageSize
}

func (p *QueryParams) Limit() int {
	return p.PageSize
}

func ParsePagination(c *gin.Context) *QueryParams {
	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		return &QueryParams{
			Page:     1,
			PageSize: 10,
		}
	}

	if params.Page <= 0 {
		params.Page = 1
	}

	if params.PageSize <= 0 || params.PageSize > 100 {
		params.PageSize = 10
	}

	return &params
}

func Paginate(query *gorm.DB, pagination *QueryParams) *gorm.DB {
	return query.Offset(pagination.Offset()).Limit(pagination.Limit())
}

func PageQuery(db *gorm.DB, model interface{}, paginationParams *QueryParams) (*PageResult, error) {
	var total int64
	countErr := db.Model(model).Count(&total).Error
	if countErr != nil {
		return nil, countErr
	}

	query := Paginate(db.Model(model), paginationParams)

	err := query.Find(model).Error
	if err != nil {
		return nil, err
	}

	paginationInfo := NewPagination(paginationParams.Page, paginationParams.PageSize, total)

	return &PageResult{
		Data:       model,
		Pagination: paginationInfo,
	}, nil
}
