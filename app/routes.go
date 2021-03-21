package app

import (
	"github.com/gorilla/mux"
	"github.com/gunturthunder/projectgolang/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router = mux.NewRouter()

}
