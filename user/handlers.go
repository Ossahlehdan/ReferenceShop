package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ReferenceShop/common"
	hadhelError "github.com/ReferenceShop/hadhelerror"
	"github.com/gorilla/mux"
)

//CreateUser create an user
func CreateUser(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "read body request failed")
		return http.StatusBadRequest, "", err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "read body request failed")
		return http.StatusBadRequest, "", err
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, err.Error())
		return http.StatusBadRequest, "", err
	}
	userCreated, err := user.SaveUser(ctx)

	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, err.Error())
		return http.StatusInternalServerError, "", err
	}
	common.OkJSON(w, userCreated, false)
	return http.StatusOK, "", nil
}

//GetUsers get all user
func GetUsers(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	user := User{}

	users, err := user.FindAllUsers(ctx)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, err.Error())
		return http.StatusInternalServerError, "", err
	}
	common.OkJSON(w, users, false)
	return http.StatusOK, "", nil
}

//GetUser get all user
func GetUser(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userid"])
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "User format not valid")
		return http.StatusBadRequest, "", err
	}
	user := User{}
	userGotten, err := user.FindUserByID(ctx, userID)
	if err != nil {
		common.KOJSON(w, http.StatusNoContent, "User not found")
		return http.StatusNoContent, "", errors.New("User  not found")
	}
	common.OkJSON(w, userGotten, false)
	return http.StatusOK, "", nil
}

//GetUserByEmailAndPwd get user by email and pwd
func GetUserByEmailAndPwd(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()
	var loginInfo LoginInfo
	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	user := User{}
	userGotten, err := user.FindUserByEmailAndPWD(ctx, loginInfo.Email, loginInfo.Password)
	if err != nil {
		common.KOJSON(w, http.StatusUnauthorized, hadhelError.ErrAPIUnauthorized.Error())
		return http.StatusUnauthorized, "", hadhelError.ErrAPIUnauthorized
	}
	common.OkJSON(w, userGotten, false)
	return http.StatusOK, "", nil
}

//UpdateUser update an user
func UpdateUser(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "")
		return http.StatusBadRequest, "", hadhelError.ErrDecoding
	}

	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, err.Error())
		return http.StatusBadRequest, "", err
	}
	updatedUser, err := user.UpdateAUser(ctx)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, err.Error())
		return http.StatusInternalServerError, "", err
	}
	common.OkJSON(w, updatedUser, false)
	return http.StatusOK, "", nil
}

//DeleteUser delete en user
func DeleteUser(w http.ResponseWriter, r *http.Request) (int, string, error) {
	ctx := r.Context()

	vars := mux.Vars(r)

	user := User{}

	userID, err := strconv.Atoi(vars["userid"])
	if err != nil {
		common.KOJSON(w, http.StatusBadRequest, "User format not valid")
		return http.StatusBadRequest, "", err
	}
	_, err = user.DeleteAUser(ctx, userID)
	if err != nil {
		common.KOJSON(w, http.StatusInternalServerError, err.Error())
		return http.StatusInternalServerError, "", err
	}
	return http.StatusOK, "", nil

}
