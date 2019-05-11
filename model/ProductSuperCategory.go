package model

type ProductSuperCategory struct {
	Id   int    `gorm:"primary_key;column:product_super_category_id" json:"id"`
	Slug string `gorm:"column:product_super_category_slug" json:"slug"`
	Name string `gorm:"column:product_super_category_name" json:"name"`
}

// func (ProductSuperCategory) TableName() string {
// 	return "asd"
// }
