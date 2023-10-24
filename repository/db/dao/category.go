package dao

import (
	"context"
	"gorm.io/gorm"
	"meetingBooking/repository/db/model"
)

type CategoryDao struct {
	*gorm.DB
}

type CategoryList struct {
	CategoryId       uint            `json:"categoryId"`
	CategoryName     string          `json:"categoryName"`
	ParentCategoryId uint            `json:"parentCategoryId"`
	CategoryLevel    uint            `json:"categoryLevel"`
	Children         []*CategoryList `json:"children"`
}

// NewCategoryDao 创建一个分类的Dto
func NewCategoryDao(ctx context.Context) *CategoryDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CategoryDao{NewDBClient(ctx)}
}

// Create 创建分类
func (dao *CategoryDao) Create(category *model.Category) (err error) {
	err = dao.DB.Model(&model.Category{}).Create(category).Error
	return
}

// FindCategoryById 根据id找到分类
func (dao *CategoryDao) FindCategoryById(id uint) (category *model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Where("category_id=?", id).
		First(&category).Error
	return
}

// UpdateCategoryById 根据id更新分类
func (dao *CategoryDao) UpdateCategoryById(id uint, category *model.Category) (err error) {
	return dao.DB.Model(&model.Category{}).Where("category_id=?", id).
		Updates(&category).Error
}

// FindCategoryAll 分类
func (dao *CategoryDao) FindCategoryAll() (category []*CategoryList, err error) {
	var categoryList []*model.Category
	err = dao.DB.Select([]string{"category_id", "category_name", "parent_category_id", "category_level"}).Find(&categoryList).Error //查询指定字段

	for _, item := range categoryList {
		categoryItem := CategoryList{
			CategoryId:       item.ID,
			CategoryName:     item.CategoryName,
			CategoryLevel:    item.CategoryLevel,
			ParentCategoryId: item.ParentCategoryID,
		}
		category = append(category, &categoryItem)
	}
	return category, nil
}

func (dao *CategoryDao) DeleteCategoryById(categoryId string) (err error) {
	err = dao.DB.Delete(&model.Category{}, categoryId).Error
	return
}
