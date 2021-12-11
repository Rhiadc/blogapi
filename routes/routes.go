package routes

import (
	"github.com/gorilla/mux"
	"github.com/rhiadc/blogapi/controllers"
)

func InitRoutes(sm *mux.Router) {
	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/users", controllers.CreateUser)
	postRouter.HandleFunc("/posts", controllers.CreatePost)
	postRouter.HandleFunc("/comments", controllers.CreateComment)

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/comments", controllers.GetComments)
	getRouter.HandleFunc("/posts", controllers.GetPosts)

}
