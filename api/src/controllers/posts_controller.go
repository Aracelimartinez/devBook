package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Adiciona uma nova publicação
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	var post *models.Post

	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	post.AuthorID = userIDToken

	err = post.Prepare()
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

	postRepo := repositories.NewPostRepository(dbConn)
	post.ID, err = postRepo.CreatePost(post)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// Traz as publicações que apareceriam no feed do usuário
func SearchPosts(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	dbConn, err := db.EstablishDbConnection()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	postRepo := repositories.NewPostRepository(dbConn)
	posts, err := postRepo.SearchPosts(userIDToken)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// Traz os dados de uma única publicação
func SearchPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["postID"], 10, 64)
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

	postRepo := repositories.NewPostRepository(dbConn)
	post, err := postRepo.SearchByID(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

// Altera os dados de uma publicação
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
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

	postRepo := repositories.NewPostRepository(dbConn)
	postInDB, err := postRepo.SearchByID(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if postInDB.AuthorID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação de outro usuário"))
		return
	}

	var post *models.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	post.AuthorID = userIDToken

	err = post.Prepare()
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	err = postRepo.UpdatePost(postID, post)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Exclui os dados de uma publicação
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
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

	postRepo := repositories.NewPostRepository(dbConn)
	postInDB, err := postRepo.SearchByID(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if postInDB.AuthorID != userIDToken {
		responses.Err(w, http.StatusForbidden, errors.New("não é possível apagar uma publicação de outro usuário"))
		return
	}

	err = postRepo.DeletePost(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Traz todas as publicações de um usuário específico
func SearchPostsByUser(w http.ResponseWriter, r *http.Request) {
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

	postRepo := repositories.NewPostRepository(dbConn)
	posts, err := postRepo.SearchByUser(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// Adiciona uma curtida na publicação
func LikePost(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
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

	postRepo := repositories.NewPostRepository(dbConn)
	err = postRepo.LikePost(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Subtrai uma curtida na publicação
func UnLikePost(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
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

	postRepo := repositories.NewPostRepository(dbConn)
	err = postRepo.UnLikePost(postID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
