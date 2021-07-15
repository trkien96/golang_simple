package route

import (
	"go_rest_api/config"
	"go_rest_api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoute() {
	Config := config.Initial()
	AppPath := Config["APP_URL"] + ":" + Config["APP_PORT"]

	route := mux.NewRouter()
	route.HandleFunc("/user/list", controller.List)
	route.HandleFunc("/user/show/{id}", controller.Show)
	route.HandleFunc("/user/create", controller.CreateForm)

	//Custom 404 page
	route.NotFoundHandler = http.HandlerFunc(controller.ShowPage404)

	log.Println("Server running: ", AppPath)
	err := http.ListenAndServe(AppPath, route)
	if err != nil {
		log.Println("Can not start server, failed: ", err)
	}
}
