package expenses

import (
	"emnaservices/webapi/internal/kernel"
	"emnaservices/webapi/utils"
	"net/http"
)

func BuildAccountService(app *kernel.Application, midd *utils.AuthMiddleware) {
	// Create our Handler
	handler := newHandler(NewService(app))

	// Register our service routes
	app.Router.HandleFunc("GET /api/expenses/find", midd.Protect(http.HandlerFunc(handler.GetExpansesData)))
	app.Router.HandleFunc("POST /api/v2/account/create", midd.Protect(http.HandlerFunc(handler.HandleCreateExpanses)))
	// app.Router.HandleFunc("POST /api/expenses/auth", handler.HandleUserLogin)

}
