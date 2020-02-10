package category

import (
	"context"
	"errors"

	"github.com/ReferenceShop/common"
	"github.com/ReferenceShop/model"
)

// FindAllCategories returns all the channels for which we have AsRun data
func FindAllCategories(ctx context.Context) ([]model.Category, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	//find all categories froom db
	var categories []model.Category
	dbConnection.Find(&categories)

	//return all categories
	return categories, nil
}

//CreateCategory  store category into db
func CreateCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Create(&category)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &category, nil
}

//FindCategoryByID  Find category by id
func FindCategoryByID(ctx context.Context, categoryID int) (*model.Category, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	var category model.Category
	dbc := dbConnection.Find(&category, categoryID)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &category, nil
}

//SaveCategory  store category into db
func SaveCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}
	dbc := dbConnection.Save(&category)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	return &category, nil
}

//DeleteCategory  delete category by id
func DeleteCategory(ctx context.Context, category model.Category) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	//delete category
	dbc := dbConnection.Delete(&category)
	if dbc.Error != nil {
		return 0, dbc.Error
	}

	return dbc.RowsAffected, nil
}
