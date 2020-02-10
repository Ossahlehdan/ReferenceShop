package product

import (
	"context"
	"errors"

	"github.com/ReferenceShop/common"
	"github.com/ReferenceShop/model"
)

// FindAllProducts returns all product
func FindAllProducts(ctx context.Context) ([]model.Product, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	//find all products froom db
	var products []model.Product
	dbConnection.Find(&products)

	//return all products
	return products, nil
}

//CreateProduct  store Product into db
func CreateProduct(ctx context.Context, product model.Product) (*model.Product, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Create(&product)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &product, nil
}

//FindProductByID  Find Product by id
func FindProductByID(ctx context.Context, ProductID int) (*model.Product, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	var product model.Product
	dbc := dbConnection.Find(&product, ProductID)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &product, nil
}

//SaveProduct  store Product into db
func SaveProduct(ctx context.Context, product model.Product) (*model.Product, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Save(&product)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &product, nil
}

//DeleteProduct  delete Product by id
func DeleteProduct(ctx context.Context, product model.Product) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	//delete Product
	dbc := dbConnection.Delete(&product)
	if dbc.Error != nil {
		return 0, dbc.Error
	}

	return dbc.RowsAffected, nil
}

// FindAllVariantes returns all variante
func FindAllVariantes(ctx context.Context) ([]model.Variante, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	//find all variantes froom db
	var variantes []model.Variante
	dbConnection.Find(&variantes)

	//return all variantes
	return variantes, nil
}

//CreateVariante  store Variante into db
func CreateVariante(ctx context.Context, variante model.Variante) (*model.Variante, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Create(&variante)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &variante, nil
}

//FindVarianteByID  Find Variante by id
func FindVarianteByID(ctx context.Context, VarianteID int) (*model.Variante, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	var variante model.Variante
	dbc := dbConnection.Find(&variante, VarianteID)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &variante, nil
}

//SaveVariante  store Variante into db
func SaveVariante(ctx context.Context, variante model.Variante) (*model.Variante, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Save(&variante)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &variante, nil
}

//DeleteVariante  delete Variante by id
func DeleteVariante(ctx context.Context, variante model.Variante) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	//delete Variante
	dbc := dbConnection.Delete(&variante)
	if dbc.Error != nil {
		return 0, dbc.Error
	}

	return dbc.RowsAffected, nil
}

// FindAllVarianteValues returns all varianteValue
func FindAllVarianteValues(ctx context.Context) ([]model.VarianteValue, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	//find all varianteValues froom db
	var varianteValues []model.VarianteValue
	dbConnection.Find(&varianteValues)

	//return all varianteValues
	return varianteValues, nil
}

//CreateVarianteValue  store VarianteValue into db
func CreateVarianteValue(ctx context.Context, varianteValue model.VarianteValue) (*model.VarianteValue, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Create(&varianteValue)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &varianteValue, nil
}

//FindVarianteValueByID  Find VarianteValue by id
func FindVarianteValueByID(ctx context.Context, VarianteValueID int) (*model.VarianteValue, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	var varianteValue model.VarianteValue
	dbc := dbConnection.Find(&varianteValue, VarianteValueID)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &varianteValue, nil
}

//SaveVarianteValue  store VarianteValue into db
func SaveVarianteValue(ctx context.Context, varianteValue model.VarianteValue) (*model.VarianteValue, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Save(&varianteValue)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &varianteValue, nil
}

//DeleteVarianteValue  delete VarianteValue by id
func DeleteVarianteValue(ctx context.Context, varianteValue model.VarianteValue) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	//delete VarianteValue
	dbc := dbConnection.Delete(&varianteValue)
	if dbc.Error != nil {
		return 0, dbc.Error
	}

	return dbc.RowsAffected, nil
}

// FindAllStocks returns all stock
func FindAllStocks(ctx context.Context) ([]model.ProductStock, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	//find all stocks froom db
	var stocks []model.ProductStock
	dbConnection.Find(&stocks)

	//return all stocks
	return stocks, nil
}

//CreateStock  store Stock into db
func CreateStock(ctx context.Context, stock model.ProductStock) (*model.ProductStock, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Create(&stock)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &stock, nil
}

//FindStockByID  Find Stock by id
func FindStockByID(ctx context.Context, StockID int) (*model.ProductStock, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	var stock model.ProductStock
	dbc := dbConnection.Find(&stock, StockID)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &stock, nil
}

//SaveStock  store Stock into db
func SaveStock(ctx context.Context, stock model.ProductStock) (*model.ProductStock, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Save(&stock)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &stock, nil
}

//DeleteStock  delete Stock by id
func DeleteStock(ctx context.Context, stock model.ProductStock) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	//delete Stock
	dbc := dbConnection.Delete(&stock)
	if dbc.Error != nil {
		return 0, dbc.Error
	}

	return dbc.RowsAffected, nil
}

// FindAllPrices returns all price
func FindAllPrices(ctx context.Context) ([]model.ProductPrice, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	//find all prices froom db
	var prices []model.ProductPrice
	dbConnection.Find(&prices)

	//return all prices
	return prices, nil
}

//CreatePrice  store Price into db
func CreatePrice(ctx context.Context, price model.ProductPrice) (*model.ProductPrice, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Create(&price)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &price, nil
}

//FindPriceByID  Find Price by id
func FindPriceByID(ctx context.Context, PriceID int) (*model.ProductPrice, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	var price model.ProductPrice
	dbc := dbConnection.Find(&price, PriceID)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &price, nil
}

//SavePrice  store Price into db
func SavePrice(ctx context.Context, price model.ProductPrice) (*model.ProductPrice, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Save(&price)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &price, nil
}

//DeletePrice  delete Price by id
func DeletePrice(ctx context.Context, price model.ProductPrice) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	//delete Price
	dbc := dbConnection.Delete(&price)
	if dbc.Error != nil {
		return 0, dbc.Error
	}

	return dbc.RowsAffected, nil
}

// FindAllImages returns all image
func FindAllImages(ctx context.Context) ([]model.ProductImage, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	//find all images froom db
	var images []model.ProductImage
	dbConnection.Find(&images)

	//return all images
	return images, nil
}

//CreateImage  store Image into db
func CreateImage(ctx context.Context, image model.ProductImage) (*model.ProductImage, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Create(&image)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &image, nil
}

//FindImageByID  Find Image by id
func FindImageByID(ctx context.Context, ImageID int) (*model.ProductImage, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	var image model.ProductImage
	dbc := dbConnection.Find(&image, ImageID)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &image, nil
}

//SaveImage  store Image into db
func SaveImage(ctx context.Context, image model.ProductImage) (*model.ProductImage, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Save(&image)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &image, nil
}

//DeleteImage  delete Image by id
func DeleteImage(ctx context.Context, image model.ProductImage) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	//delete Image
	dbc := dbConnection.Delete(&image)
	if dbc.Error != nil {
		return 0, dbc.Error
	}

	return dbc.RowsAffected, nil
}
