package endpoints

import (
	"api/src/controllers"
	"net/http"
)

var userEndpoints = []Endopoints{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     controllers.SearchUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/unfollow",
		Method:       http.MethodDelete,
		Function:     controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/followers",
		Method:       http.MethodGet,
		Function:     controllers.SearchUserFollowers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/following",
		Method:       http.MethodGet,
		Function:     controllers.SearchUsersFollowed,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/update-password",
		Method:       http.MethodPost,
		Function:     controllers.UpdateUserPassword,
		AuthRequired: true,
	},
}
