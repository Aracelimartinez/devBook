package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"net/http"
)

// Responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	err := json.NewDecoder(r.Body).Decode(&user)
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
	userDB, err := userRepo.SearchUserByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPassword(userDB.Password, user.Password)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userDB.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))
}

