package model

import (
	gorm "github.com/jinzhu/gorm"
)

//Category instance representation
type Category struct {
	gorm.Model
	CategoryCode string     `gorm:"unique;not null" json:"categoryCode"`
	CategoryName string     `gorm:"unique;not null" json:"categoryName"`
	IsActive     bool       `json:"isActive"`
	Description  string     `json:"description"`
	Parent       *Category  `gorm:"Foreignkey:ParentID" json:"parent"`
	ParentID     int        `json:"parentID"`
	Products     []*Product `gorm:"many2many:products"`
	HasStock     bool       `json:"hasStock"`
}

//Product represente product data
type Product struct {
	gorm.Model
	ProductCode string      `gorm:"unique;not null" json:"productCode"`
	ProductName string      `gorm:"unique;not null" json:"productName"`
	IsActive    bool        ` json:"isActive"`
	Description string      ` json:"description"`
	Categories  []*Category `gorm:"many2many:categories"`
}

//Variante represent variante data
type Variante struct {
	gorm.Model
	VarianteName string `gorm:"unique;not null" json:"varianteName"`
	VarinateType string `gorm:"not null" json:"varianteType"`
	IsActive     bool   `json:"isActive"`
}

//VarianteValue represente varaiante valu data
type VarianteValue struct {
	gorm.Model
	Value      string    `json:"value"`
	ColorHTML  string    `json:"colorHtml"`
	Variante   *Variante `gorm:"Foreignkey:ParentID" json:"variante"`
	VarianteID int       `json:"varianteId"`
}

//ProductStock répresent the product inventory
type ProductStock struct {
	gorm.Model
	InitialStock int16    `json:"intialStock"`
	Stock        int16    `json:"stock"`
	Product      *Product `gorm:"Foreignkey:ProductID" json:"product"`
	ProductID    int      `json:"productId"`
	Variante     *Product `gorm:"Foreignkey:VarianteID" json:"variante"`
	VarianteID   int      `json:"varianteId"`
}

//ProductPrice répresent the product price
type ProductPrice struct {
	gorm.Model
	PriceCode  float64  `json:"priceCode"`
	Price      float64  `json:"price"`
	IsDiscount bool     `json:"isDiscount"`
	IsPercent  bool     `json:"isPercent"`
	Product    *Product `gorm:"Foreignkey:ParentID" json:"product"`
	ProductID  int      `json:"productId"`
	Variante   *Product `gorm:"Foreignkey:VarianteID" json:"variante"`
	VarianteID int      `json:"varianteId"`
}

//ProductImage répresent the product image
type ProductImage struct {
	gorm.Model
	Name       float64  `json:"priceName"`
	Path       float64  `json:"path"`
	IsDefault  bool     `json:"isDefault"`
	Product    *Product `gorm:"Foreignkey:ParentID" json:"product"`
	ProductID  int      `json:"productId"`
	Variante   *Product `gorm:"Foreignkey:VarianteID" json:"variante"`
	VarianteID int      `json:"varianteId"`
}
