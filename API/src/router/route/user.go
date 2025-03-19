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
		AuthRequested: true,
	},
	{
		Uri:           "/user",
		Method:        http.MethodPost,
		Func:          controllers.CreateUser,
		AuthRequested: true,
	},
	{
		Uri:           "/user/{userId}",
		Method:        http.MethodGet,
		Func:          controllers.GetUser,
		AuthRequested: true,
	},
	{
		Uri:           "/user/{userId}",
		Method:        http.MethodPut,
		Func:          controllers.UpdateUser,
		AuthRequested: true,
	},
	{
		Uri:           "/user/{userId}",
		Method:        http.MethodDelete,
		Func:          controllers.DeleteUser,
		AuthRequested: true,
	},
}
