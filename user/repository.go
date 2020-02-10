package user

import (
	"context"
	"errors"
	"log"

	"github.com/Hadhel.ReferenceShop/common"
	"github.com/jinzhu/gorm"
	bcrypt "golang.org/x/crypto/bcrypt"
)

//SaveUser  create user
func (u *User) SaveUser(ctx context.Context) (*User, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}

	err := dbConnection.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//FindAllUsers retrieve 100 first users
func (u *User) FindAllUsers(ctx context.Context) (*[]User, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}

	users := []User{}
	err := dbConnection.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

//FindUserByID  get user by id
func (u *User) FindUserByID(ctx context.Context, userID int) (*User, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}

	err := dbConnection.Debug().Model(User{}).Where("id = ?", userID).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

//FindUserByEmailAndPWD  get user by id
func (u *User) FindUserByEmailAndPWD(ctx context.Context, email, password string) (*User, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}

	err := dbConnection.Debug().Model(User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	err = VerifyPassword(u.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

//UpdateAUser update user
func (u *User) UpdateAUser(ctx context.Context) (*User, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return nil, errors.New("Could not get database connection from context")
	}

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db := dbConnection.Debug().Model(&User{}).Save(u)

	if db.Error != nil {
		return &User{}, db.Error
	}

	return u, nil
}

//DeleteAUser delete user
func (u *User) DeleteAUser(ctx context.Context, userID int) (int64, error) {
	//Get db connexion
	dbConnection, ok := common.GetDatabaseConnection(ctx)
	if !ok {
		return 0, errors.New("Could not get database connection from context")
	}
	db := dbConnection.Debug().Model(&User{}).Where("id = ?", userID).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
