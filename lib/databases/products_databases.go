package databases

import (
	"final-project/config"
	"final-project/models"
)

var get_product []models.GetProduct

// function database untuk membuat data product baru
func CreateProduct(product *models.Products) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// function database untuk menampilkan data semua product
func GetAllProduct() (interface{}, error) {
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Find(&get_product)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_product, nil
}

// function database untuk menampilkan data product by id
func GetProductById(id int) (interface{}, error) {
	var get_product_by_id models.GetProduct
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL AND products.id = ?", id).Find(&get_product_by_id)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_product_by_id, nil
}
