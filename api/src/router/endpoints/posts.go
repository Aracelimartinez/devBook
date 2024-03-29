package endpoints

import (
	"api/src/controllers"
	"net/http"
)

var postsEndpoints = []Endopoints{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		Function:     controllers.SearchPosts,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodGet,
		Function:     controllers.SearchPost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/posts",
		Method:       http.MethodGet,
		Function:     controllers.SearchPostsByUser,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnLikePost,
		AuthRequired: true,
	},
}
