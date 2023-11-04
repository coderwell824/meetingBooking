package services

import (
	"context"
	"errors"
	"meetingBooking/config"
	"meetingBooking/consts"
	"meetingBooking/repository/db/dao"
	"meetingBooking/repository/db/model"
	"meetingBooking/utils"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	ID            uint    `form:"id" json:"id"`
	Name          string  `form:"name" json:"name"`
	CategoryID    uint    `form:"categoryId" json:"categoryId"`
	Title         string  `form:"title" json:"title"`
	Info          string  `form:"info" json:"info"`
	ImgPath       string  `form:"imgPath" json:"imgPath" `
	Price         float64 `form:"price" json:"price" `
	DiscountPrice float64 `form:"discountPrice" json:"discountPrice"`
	OnSale        bool    `form:"onSale" json:"onSale"`
	Num           uint    `form:"num" json:"num"`
	PageNum       int     `form:"pageNum"`
	PageSize      int     `form:"pageSize"`
}

func (service *ProductService) Create(ctx context.Context, files []*multipart.FileHeader) (resp string, err error) {

	var path string
	u, err := utils.GetUserInfo(ctx)
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)
	temp, _ := files[0].Open()
	if config.UploadModel == consts.UploadModelLocal {
		path, err = utils.UploadAvatarToLocalStatic(temp, user.ID, user.Username)
	} else {
		path, err = utils.UploadImageToQiQiu(temp, files[0].Size)
	}
	if err != nil {
		err = errors.New("上传图片失败")
		return
	}

	product := &model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryID,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DisCountPrice: service.DiscountPrice,
		Num:           service.Num,
		OnSale:        true,
		BossID:        u.Id,
		BossName:      user.Username,
		BossAvatar:    user.AvatarUrl,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.Create(product)
	if err != nil {
		err = errors.New("添加商品失败")
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ := file.Open()
		if config.UploadModel == consts.UploadModelLocal {
			path, err = utils.UploadAvatarToLocalStatic(tmp, user.ID, service.Name+num)
		} else {
			path, err = utils.UploadImageToQiQiu(tmp, file.Size)
		}
		if err != nil {
			err = errors.New("上传文件失败")
			return
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			err = errors.New("创建商品图片失败")
			return
		}
		wg.Done()
	}

	wg.Wait()

	return "创建商品成功", nil
}

func (service *ProductService) List(ctx context.Context) (resp interface{}, total int64, err error) {
	var products []*model.Product
	//TODO： pageNum和pageSize的校验

	condition := make(map[string]interface{})
	if service.CategoryID != 0 {
		condition["category_id"] = service.CategoryID
	}
	productDao := dao.NewProductDao(ctx)
	total, err = productDao.CountProductByCondition(condition)
	if err != nil {
		err = errors.New("获取商品列表失败")
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.PageNum, service.PageSize)
		wg.Done()
	}()
	wg.Wait()
	return products, total, nil
}

func (service *ProductService) Update(ctx context.Context, pId string) (resp string, err error) {
	productDao := dao.NewProductDao(ctx)
	productId, _ := strconv.Atoi(pId)
	newProduct := &model.Product{
		Name:          service.Name,
		CategoryID:    uint(service.CategoryID),
		Title:         service.Title,
		Info:          service.Info,
		Price:         service.Price,
		DisCountPrice: service.DiscountPrice,
		OnSale:        service.OnSale,
		Num:           service.Num,
	}
	err = productDao.UpdateProductById(uint(productId), newProduct)
	if err != nil {
		err = errors.New("商品更新失败")
		return
	}
	return "更新商品成功", nil
}
