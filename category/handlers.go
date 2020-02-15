package category

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ReferenceShop/common"
	hadhelError "github.com/ReferenceShop/hadhelerror"
	"github.com/ReferenceShop/model"
	"github.com/gorilla/mux"
)

//GetCategories get Category
func GetCategories(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Get all categories from db
	categories, err := FindAllCategories(ctx)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	if len(categories) == 0 {
		return http.StatusNoContent, "", errors.New("No Categories found")
	}
	common.OkJSON(w, categories, false)
	return http.StatusOK, "", nil
}

//AddCategory  get request json body, create and save in db
func AddCategory(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	//Decode json body request
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}
	fmt.Println(category)
	//Store into db
	newCategory, err := CreateCategory(ctx, category)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Store category in db failed")
		return http.StatusInternalServerError, "", errors.New("Store category in db failed")
	}

	common.OkJSON(w, newCategory, false)
	return http.StatusOK, "", nil
}

//GetCayegoryByID  find category by id
func GetCayegoryByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	categoryID, err := strconv.Atoi(vars["categoryId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Category Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Category Id format not valid")
	}
	//Find category by id
	category, err := FindCategoryByID(ctx, categoryID)

	fmt.Println(category)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Category not found")
		return http.StatusNoContent, "", errors.New("category  not found")
	}

	common.OkJSON(w, category, false)
	return http.StatusOK, "", nil
}

//UpdateCategory  update Category
func UpdateCategory(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	//Decode json body request
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	//Store into db
	UpdatedCategory, err := SaveCategory(ctx, category)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Save category in db failed")
		return http.StatusInternalServerError, "", errors.New("Save category in db failed")
	}

	common.OkJSON(w, UpdatedCategory, false)
	return http.StatusOK, "", nil
}

//DeleteCategoryByID  find category by id
func DeleteCategoryByID(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	categoryID, err := strconv.Atoi(vars["categoryId"])
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, "Category Id format not valid")
		return http.StatusInternalServerError, "", errors.New("Category Id format not valid")
	}
	//find category by id
	category, err := FindCategoryByID(ctx, categoryID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Category not found")
		return http.StatusNoContent, "", errors.New("category  not found")
	}

	//Find category by id
	_, err = DeleteCategory(ctx, *category)

	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "Category not found")
		return http.StatusNoContent, "", errors.New("category  not found")
	}

	return http.StatusOK, "", nil
}
