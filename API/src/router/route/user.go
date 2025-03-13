package route

import (
	"example.com/m/v2/API/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		Uri:           "/user",
		Method:        http.MethodGet,
		Func:          controllers.ListUsers,
		AuthRequested: false,
	},
	{
		Uri:           "/user",
		Method:        http.MethodPost,
		Func:          controllers.CreateUser,
		AuthRequested: false,
	},
	{
		Uri:           "/user/{userId}",
		Method:        http.MethodGet,
		Func:          controllers.GetUser,
		AuthRequested: false,
	},
	{
		Uri:           "/user/{userId}",
		Method:        http.MethodPut,
		Func:          controllers.UpdateUser,
		AuthRequested: false,
	},
	{
		Uri:           "/user/{userId}",
		Method:        http.MethodDelete,
		Func:          controllers.DeleteUser,
		AuthRequested: false,
	},
}
