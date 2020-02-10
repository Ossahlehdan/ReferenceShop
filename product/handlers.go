package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hadhel.ReferenceShop/common"
	hadhelError "github.com/Hadhel.ReferenceShop/hadhelerror"
	"github.com/Hadhel.ReferenceShop/model"
	"github.com/gorilla/mux"
)

//GetProducts get Product
func GetProducts(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Get all Products from db
	products, err := FindAllProducts(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	common.OkJSON(w, products, false)

	return http.StatusOK, "", nil
}

//AddProduct  get request json body, create and save in db
func AddProduct(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Decode json body request
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}
	fmt.Println(product)
	//Store into db
	newProduct, err := CreateProduct(ctx, product)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Store Product in db failed")
		return http.StatusInternalServerError, "", errors.New("Store Product in db failed")
	}

	common.OkJSON(w, newProduct, false)
	return http.StatusOK, "", nil
}

//GetProductByID  find Product by id
func GetProductByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productID, err := strconv.Atoi(vars["ProductId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Product Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Product Id format not valid")
	}
	//Find Product by id
	product, err := FindProductByID(ctx, productID)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Product not found")
		return http.StatusNoContent, "", errors.New("Product  not found")
	}

	common.OkJSON(w, product, false)
	return http.StatusOK, "", nil
}

//UpdateProduct  update Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	//Decode json body request
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	//Store into db
	UpdatedProduct, err := SaveProduct(ctx, product)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Save Product in db failed")
		return http.StatusInternalServerError, "", errors.New("Save Product in db failed")
	}

	common.OkJSON(w, UpdatedProduct, false)
	return http.StatusOK, "", nil
}

//DeleteProductByID  find Product by id
func DeleteProductByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productID, err := strconv.Atoi(vars["ProductId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Product Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Product Id format not valid")
	}
	//find Product by id
	product, err := FindProductByID(ctx, productID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Product not found")
		return http.StatusNoContent, "", errors.New("Product  not found")
	}

	//Find Product by id
	_, err = DeleteProduct(ctx, *product)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Product not found")
		return http.StatusNoContent, "", errors.New("Product  not found")
	}

	return http.StatusOK, "", nil
}

//GetProductVariantes get Product
func GetProductVariantes(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Get all Products from db
	products, err := FindAllVariantes(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	common.OkJSON(w, products, false)

	return http.StatusOK, "", nil
}

//AddProductVariante  get request json body, create and save in db
func AddProductVariante(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Decode json body request
	var variante model.Variante
	err := json.NewDecoder(r.Body).Decode(&variante)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}
	fmt.Println(variante)
	//Store into db
	newvariante, err := CreateVariante(ctx, variante)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Store Variante in db failed")
		return http.StatusInternalServerError, "", errors.New("Store Variante in db failed")
	}

	common.OkJSON(w, newvariante, false)
	return http.StatusOK, "", nil
}

//GetProductVarianteByID  find Product by id
func GetProductVarianteByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	varianteID, err := strconv.Atoi(vars["varianteId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Variante Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Variante Id format not valid")
	}
	//Find Variante by id
	variante, err := FindVarianteByID(ctx, varianteID)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Variante not found")
		return http.StatusNoContent, "", errors.New("Variante  not found")
	}

	common.OkJSON(w, variante, false)
	return http.StatusOK, "", nil
}

//UpdateProductVariante  update Product
func UpdateProductVariante(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	//Decode json body request
	var variante model.Variante
	err := json.NewDecoder(r.Body).Decode(&variante)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	//Store into db
	UpdatedVariante, err := SaveVariante(ctx, variante)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Save variante in db failed")
		return http.StatusInternalServerError, "", errors.New("Save variante in db failed")
	}

	common.OkJSON(w, UpdatedVariante, false)
	return http.StatusOK, "", nil
}

//DeleteProductVarianteByID  find Product by id
func DeleteProductVarianteByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	varianteID, err := strconv.Atoi(vars["varianteId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Variante Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Variante Id format not valid")
	}
	//find variante by id
	variante, err := FindVarianteByID(ctx, varianteID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Variante not found")
		return http.StatusNoContent, "", errors.New("Variante  not found")
	}

	//Find variante by id
	_, err = DeleteVariante(ctx, *variante)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Delete variante failed")
		return http.StatusNoContent, "", errors.New("Delete variante failed")
	}

	return http.StatusOK, "", nil
}

//GetProductVarianteValues get Product
func GetProductVarianteValues(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Get all Products from db
	products, err := FindAllVarianteValues(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	common.OkJSON(w, products, false)

	return http.StatusOK, "", nil
}

//AddProductVarianteValue  get request json body, create and save in db
func AddProductVarianteValue(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Decode json body request
	var varianteValue model.VarianteValue
	err := json.NewDecoder(r.Body).Decode(&varianteValue)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}
	fmt.Println(varianteValue)
	//Store into db
	newVarianteValue, err := CreateVarianteValue(ctx, varianteValue)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Store Variante value in db failed")
		return http.StatusInternalServerError, "", errors.New("Store Variante value in db failed")
	}

	common.OkJSON(w, newVarianteValue, false)
	return http.StatusOK, "", nil
}

//GetProductVarianteValueByID  find Product by id
func GetProductVarianteValueByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	varianteValueID, err := strconv.Atoi(vars["varianteValueId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Variante Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Variante Id format not valid")
	}
	//Find Variante by id
	varianteValue, err := FindVarianteValueByID(ctx, varianteValueID)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Variante not found")
		return http.StatusNoContent, "", errors.New("Variante  not found")
	}

	common.OkJSON(w, varianteValue, false)
	return http.StatusOK, "", nil
}

//UpdateProductVarianteValue  update Product
func UpdateProductVarianteValue(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	//Decode json body request
	var varianteValue model.VarianteValue
	err := json.NewDecoder(r.Body).Decode(&varianteValue)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	//Store into db
	UpdatedVarianteValue, err := SaveVarianteValue(ctx, varianteValue)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Save varianteValue in db failed")
		return http.StatusInternalServerError, "", errors.New("Save varianteValue in db failed")
	}

	common.OkJSON(w, UpdatedVarianteValue, false)
	return http.StatusOK, "", nil
}

//DeleteProductVarianteValueByID  find Product by id
func DeleteProductVarianteValueByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	varianteValueID, err := strconv.Atoi(vars["varianteValueId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Variante value Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Variante value Id format not valid")
	}
	//find varianteValue by id
	varianteValue, err := FindVarianteValueByID(ctx, varianteValueID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Variante value not found")
		return http.StatusNoContent, "", errors.New("Variante value not found")
	}

	//Find varianteValue by id
	_, err = DeleteVarianteValue(ctx, *varianteValue)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Delete varianteValue failed")
		return http.StatusNoContent, "", errors.New("Delete varianteValue failed")
	}

	return http.StatusOK, "", nil
}

//GetProductStocks get ProductStock
func GetProductStocks(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Get all ProductStocks from db
	productStocks, err := FindAllStocks(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	common.OkJSON(w, productStocks, false)

	return http.StatusOK, "", nil
}

//AddProductStock  get request json body, create and save in db
func AddProductStock(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Decode json body request
	var productStock model.ProductStock
	err := json.NewDecoder(r.Body).Decode(&productStock)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}
	fmt.Println(productStock)
	//Store into db
	newProductStock, err := CreateStock(ctx, productStock)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Store ProductStock in db failed")
		return http.StatusInternalServerError, "", errors.New("Store ProductStock in db failed")
	}

	common.OkJSON(w, newProductStock, false)
	return http.StatusOK, "", nil
}

//GetProductStockByID  find ProductStock by id
func GetProductStockByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productStockID, err := strconv.Atoi(vars["ProductStockId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "ProductStock Id format not valid")
		return http.StatusInternalServerError, "", errors.New("ProductStock Id format not valid")
	}
	//Find ProductStock by id
	productStock, err := FindStockByID(ctx, productStockID)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductStock not found")
		return http.StatusNoContent, "", errors.New("ProductStock  not found")
	}

	common.OkJSON(w, productStock, false)
	return http.StatusOK, "", nil
}

//UpdateProductStock  update ProductStock
func UpdateProductStock(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	//Decode json body request
	var productStock model.ProductStock
	err := json.NewDecoder(r.Body).Decode(&productStock)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	//Store into db
	UpdatedProductStock, err := SaveStock(ctx, productStock)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Save ProductStock in db failed")
		return http.StatusInternalServerError, "", errors.New("Save ProductStock in db failed")
	}

	common.OkJSON(w, UpdatedProductStock, false)
	return http.StatusOK, "", nil
}

//DeleteProductStockByID  find ProductStock by id
func DeleteProductStockByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productStockID, err := strconv.Atoi(vars["ProductStockId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "ProductStock Id format not valid")
		return http.StatusInternalServerError, "", errors.New("ProductStock Id format not valid")
	}
	//find ProductStock by id
	productStock, err := FindStockByID(ctx, productStockID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductStock not found")
		return http.StatusNoContent, "", errors.New("ProductStock  not found")
	}

	//Find ProductStock by id
	_, err = DeleteStock(ctx, *productStock)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductStock not found")
		return http.StatusNoContent, "", errors.New("ProductStock  not found")
	}

	return http.StatusOK, "", nil
}

//GetProductPrices get ProductPrice
func GetProductPrices(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Get all ProductPrices from db
	productPrices, err := FindAllStocks(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	common.OkJSON(w, productPrices, false)

	return http.StatusOK, "", nil
}

//AddProductPrice  get request json body, create and save in db
func AddProductPrice(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Decode json body request
	var productPrice model.ProductPrice
	err := json.NewDecoder(r.Body).Decode(&productPrice)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}
	fmt.Println(productPrice)
	//Store into db
	newProductPrice, err := CreatePrice(ctx, productPrice)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Store ProductPrice in db failed")
		return http.StatusInternalServerError, "", errors.New("Store ProductPrice in db failed")
	}

	common.OkJSON(w, newProductPrice, false)
	return http.StatusOK, "", nil
}

//GetProductPriceByID  find ProductPrice by id
func GetProductPriceByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productPriceID, err := strconv.Atoi(vars["ProductPriceId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "ProductPrice Id format not valid")
		return http.StatusInternalServerError, "", errors.New("ProductPrice Id format not valid")
	}
	//Find ProductPrice by id
	productPrice, err := FindStockByID(ctx, productPriceID)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductPrice not found")
		return http.StatusNoContent, "", errors.New("ProductPrice  not found")
	}

	common.OkJSON(w, productPrice, false)
	return http.StatusOK, "", nil
}

//UpdateProductPrice  update ProductPrice
func UpdateProductPrice(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	//Decode json body request
	var productPrice model.ProductPrice
	err := json.NewDecoder(r.Body).Decode(&productPrice)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	//Store into db
	UpdatedProductPrice, err := SavePrice(ctx, productPrice)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Save ProductPrice in db failed")
		return http.StatusInternalServerError, "", errors.New("Save ProductPrice in db failed")
	}

	common.OkJSON(w, UpdatedProductPrice, false)
	return http.StatusOK, "", nil
}

//DeleteProductPriceByID  find ProductPrice by id
func DeleteProductPriceByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productPriceID, err := strconv.Atoi(vars["ProductPriceId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "ProductPrice Id format not valid")
		return http.StatusInternalServerError, "", errors.New("ProductPrice Id format not valid")
	}
	//find ProductPrice by id
	productPrice, err := FindPriceByID(ctx, productPriceID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductPrice not found")
		return http.StatusNoContent, "", errors.New("ProductPrice  not found")
	}

	//Find ProductPrice by id
	_, err = DeletePrice(ctx, *productPrice)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductPrice not found")
		return http.StatusNoContent, "", errors.New("ProductPrice  not found")
	}

	return http.StatusOK, "", nil
}

//GetProductImages get ProductImage
func GetProductImages(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Get all ProductImages from db
	productImages, err := FindAllImages(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	common.OkJSON(w, productImages, false)

	return http.StatusOK, "", nil
}

//AddProductImage  get request json body, create and save in db
func AddProductImage(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Decode json body request
	var productImage model.ProductImage
	err := json.NewDecoder(r.Body).Decode(&productImage)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}
	fmt.Println(productImage)
	//Store into db
	newProductImage, err := CreateImage(ctx, productImage)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Store ProductImage in db failed")
		return http.StatusInternalServerError, "", errors.New("Store ProductImage in db failed")
	}

	common.OkJSON(w, newProductImage, false)
	return http.StatusOK, "", nil
}

//GetProductImageByID  find ProductImage by id
func GetProductImageByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productImageID, err := strconv.Atoi(vars["ProductImageId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "ProductImage Id format not valid")
		return http.StatusInternalServerError, "", errors.New("ProductImage Id format not valid")
	}
	//Find ProductImage by id
	productImage, err := FindImageByID(ctx, productImageID)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductImage not found")
		return http.StatusNoContent, "", errors.New("ProductImage  not found")
	}

	common.OkJSON(w, productImage, false)
	return http.StatusOK, "", nil
}

//UpdateProductImage  update ProductImage
func UpdateProductImage(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	//Decode json body request
	var productImage model.ProductImage
	err := json.NewDecoder(r.Body).Decode(&productImage)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	//Store into db
	UpdatedProductImage, err := SaveImage(ctx, productImage)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Save ProductImage in db failed")
		return http.StatusInternalServerError, "", errors.New("Save ProductImage in db failed")
	}

	common.OkJSON(w, UpdatedProductImage, false)
	return http.StatusOK, "", nil
}

//DeleteProductImageByID  find ProductImage by id
func DeleteProductImageByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	productImageID, err := strconv.Atoi(vars["ProductImageId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "ProductImage Id format not valid")
		return http.StatusInternalServerError, "", errors.New("ProductImage Id format not valid")
	}
	//find ProductImage by id
	productImage, err := FindImageByID(ctx, productImageID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductImage not found")
		return http.StatusNoContent, "", errors.New("ProductImage  not found")
	}

	//Find ProductImage by id
	_, err = DeleteImage(ctx, *productImage)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "ProductImage not found")
		return http.StatusNoContent, "", errors.New("ProductImage  not found")
	}

	return http.StatusOK, "", nil
}
