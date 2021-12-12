package routes

import (
	"github.com/gorilla/mux"
	"github.com/rhiadc/blogapi/controllers"
)

func InitRoutes(sm *mux.Router) {
	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/users", controllers.CreateUser)
	postRouter.HandleFunc("/posts", controllers.CreatePost)
	postRouter.HandleFunc("/posts/{id}/comments", controllers.CreateComment)

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/posts/{id}/comments", controllers.GetComments)
	getRouter.HandleFunc("/posts/{id}", controllers.GetPost)
	getRouter.HandleFunc("/posts", controllers.GetPosts)
	getRouter.HandleFunc("/users", controllers.GetUsers)
	getRouter.HandleFunc("/users/{id}", controllers.GetUser)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/users/{id}", controllers.DeleteUser)
	deleteRouter.HandleFunc("/posts/{id}/comments/{comment_id}", controllers.DeleteComment)
	deleteRouter.HandleFunc("/posts/{id}", controllers.DeletePost)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/users/{id}", controllers.UpdateUser)
	putRouter.HandleFunc("/post/{id}", controllers.UpdatePost)
	putRouter.HandleFunc("/posts/{id}/comments/{comment_id}", controllers.UpdateComment)
}
