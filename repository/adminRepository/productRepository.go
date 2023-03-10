package adminRepository

import (
	"Tani-Desa/dto/adminDto"
	"Tani-Desa/model"
	"time"
)

func (u *adminRepository) CreateProduct(payloads adminDto.ProductRequest, link string) error {

	if err := u.db.Create(&model.Products{
		AdminId:     payloads.AdminID,
		Name:        payloads.Name,
		Image:       link,
		Type:        payloads.Type,
		Qty:         payloads.Qty,
		Price:       payloads.Price,
		Description: payloads.Description,
		Address:     payloads.Address,
		Owner:       payloads.Owner,
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (u *adminRepository) GetAllProducts(adminId uint) ([]adminDto.ProductDTO, error) {
	var products []adminDto.ProductDTO

	if err := u.db.Model(&model.Products{}).Select("products.*,admins.username as admin_name").
		Joins("join admins on admins.admin_id = products.admin_id").
		Where("products.admin_id=?", adminId).Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (u *adminRepository) UpdateImageProduct(productId uint, link string) error {
	if err := u.db.Model(&model.Products{}).Where("product_id=?", productId).Updates(&model.Products{
		Image:     link,
		UpdatedAT: time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil

}

func (u *adminRepository) GetProductById(adminId uint, productId uint) (adminDto.ProductDTO, error) {
	product := adminDto.ProductDTO{}
	if err := u.db.Model(&model.Products{}).Select("products.*,admins.username as admin_name").
		Joins("join admins on admins.admin_id =products.admin_id").Where("products.admin_id=? && product_id =?", adminId, productId).Find(&product).
		Error; err != nil {
		return product, err
	}
	return product, nil
}

func (u *adminRepository) DeleteProductById(adminId uint, productId uint) error {
	if err := u.db.Where("products.admin_id=? && product_id =?", adminId, productId).Delete(&model.Products{}).Error; err != nil {
		return err
	}
	return nil
}
