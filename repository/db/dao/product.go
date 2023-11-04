package dao

import (
	"context"
	"gorm.io/gorm"
	"meetingBooking/repository/db/model"
)

type ProductDao struct {
	*gorm.DB
}

// NewProductDao 创建一个商品的Dto
func NewProductDao(ctx context.Context) *ProductDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// Create 创建商品
func (dao *ProductDao) Create(product *model.Product) (err error) {
	err = dao.DB.Model(&model.Product{}).Create(&product).Error
	return
}

// CountProductByCondition 获取商品的数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// ListProductByCondition 获取商品列表
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, pageNum, pageSize int) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&products).Error
	return
}

// UpdateProductById 更新商品
func (dao *ProductDao) UpdateProductById(pId uint, newProduct *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("product_id=?", pId).Updates(&newProduct).Error
}
