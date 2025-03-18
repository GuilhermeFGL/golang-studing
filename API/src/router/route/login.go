package route

import (
	"example.com/m/v2/API/src/controllers"
	"net/http"
)

var LoginRoutes = []Route{
	{
		Uri:           "/login",
		Method:        http.MethodPost,
		Func:          controllers.Login,
		AuthRequested: false,
	},
}
