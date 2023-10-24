package services

import (
	"context"
	"errors"
	"log"
	"meetingBooking/repository/db/dao"
	"meetingBooking/repository/db/model"
	"meetingBooking/utils"
)

type CategoryService struct {
	CategoryName     string `json:"categoryName" binding:"required" msg:"分类名称不能为空"`
	ParentCategoryId uint   `json:"parentCategoryId"`
}

type UpdateCategoryService struct {
	CategoryId   uint   `json:"categoryId" binding:"required" msg:"分类Id不能为空"`
	CategoryName string `json:"categoryName" binding:"required" msg:"分类名称不能为空"`
}

func (s *CategoryService) Create(ctx context.Context, category CategoryService) (response interface{}, err error) {
	categoryDao := dao.NewCategoryDao(ctx)
	categories := model.Category{
		CategoryName: category.CategoryName,
	}
	if category.ParentCategoryId == 0 {
		categories.CategoryLevel = 1
	} else {
		categories.ParentCategoryID = category.ParentCategoryId
		findCategory, err := categoryDao.FindCategoryById(category.ParentCategoryId)
		if err != nil {
			log.Println(err)
			return "", errors.New("新增分类失败")
		}
		if findCategory.CategoryLevel == 1 {
			categories.CategoryLevel = 2
		} else {
			categories.CategoryLevel = 3
		}

	}

	err = categoryDao.Create(&categories)
	if err != nil {
		log.Println(err)
		return "", errors.New("新增分类失败")
	}
	return "新增分类成功", nil
}

func (s *UpdateCategoryService) Update(ctx context.Context, category UpdateCategoryService) (response interface{}, err error) {
	categoryDao := dao.NewCategoryDao(ctx)
	findCategory, err := categoryDao.FindCategoryById(category.CategoryId)
	if err != nil {
		log.Println(err)
		return "", errors.New("分类id不存在")
	}
	findCategory.CategoryName = category.CategoryName
	err = categoryDao.UpdateCategoryById(category.CategoryId, findCategory)
	if err != nil {
		log.Println(err)
		return "", errors.New("更新分类名称失败")
	}
	return "更新分类名称成功", nil
}

func (s *CategoryService) List(ctx context.Context) (response interface{}, err error) {
	categoryDao := dao.NewCategoryDao(ctx)
	findCategoryList, _ := categoryDao.FindCategoryAll()
	categoryList := utils.BuildCategoryTree(findCategoryList)
	return categoryList, nil
}
func (s *CategoryService) Delete(ctx context.Context, categoryId string) (response interface{}, err error) {
	categoryDao := dao.NewCategoryDao(ctx)
	err = categoryDao.DeleteCategoryById(categoryId)
	if err != nil {
		return "", err
	}
	return "删除分类成功", nil
}
