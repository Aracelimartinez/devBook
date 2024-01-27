package endpoints

import (
	"api/src/controllers"
	"net/http"
)

var loginEndpoint = Endopoints{
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	AuthRequired: false,
}
