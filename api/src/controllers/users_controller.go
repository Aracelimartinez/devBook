package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Cria um novo usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser *models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	err = newUser.Prepare("register")
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	userID, err := userRepo.CreateUser(newUser)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, userID)
}

// Obtem usuários que atendem um filtro de nome ou nickname
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := r.URL.Query().Get("user")

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	users, err := userRepo.SearchUsers(nameOrNick)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// Busca um usuário
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	user, err := userRepo.SearchUserByID(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// Altera as informações de um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("você não está autorizado para atualizar este usuário"))
		return
	}

	var updatedUser *models.User

	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	err = updatedUser.Prepare("edit")
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	err = userRepo.UpdateUser(userID, updatedUser)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// Exclui as informações de um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("você não está autorizado para apagar este usuário"))
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	err = userRepo.DeleteUser(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// Permite que um usuário siga outro
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Err(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	err = userRepo.Follow(userID, followerID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Permite que um usuário pare de seguir outro
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Err(w, http.StatusForbidden, errors.New("não é possível parar de seguir você mesmo"))
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	err = userRepo.Unfollow(userID, followerID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Permite vizualizar os seguidores de um usuário
func SearchUserFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	followers, err := userRepo.SearchFollowers(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// Traz todos os usuários que um determinado usuário está seguindo
func SearchUsersFollowed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	usersFollowed, err := userRepo.SearchUsersFollowed(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, usersFollowed)
}

// Altera a senha do usuário
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("você não está autorizado para atualizar a senha deste usuário"))
		return
	}

	var password *models.Password

	err = json.NewDecoder(r.Body).Decode(&password)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	userRepo := repositories.NewUsersRepository(dbConn)
	passwordInDB, err := userRepo.SearchPassword(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPassword(passwordInDB, password.Current)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, errors.New("A senha atual está incorreta"))
		return
	}

	hashPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	err = userRepo.UpdatePassword(userID, string(hashPassword))
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
