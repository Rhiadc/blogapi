package routes

import (
	"github.com/gorilla/mux"
	"github.com/rhiadc/blogapi/controllers"
	"github.com/rhiadc/blogapi/middlewares"
)

func InitRoutes(sm *mux.Router) {
	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/posts", controllers.CreatePost)
	postRouter.HandleFunc("/posts/{id}/comments", controllers.CreateComment)

	postRouter.Use(middlewares.TokenAuthMiddleware)

	loginRouter := sm.Methods("POST").Subrouter()
	loginRouter.HandleFunc("/login", controllers.Login)
	loginRouter.HandleFunc("/users", controllers.CreateUser)
	loginRouter.HandleFunc("/logout", controllers.Logout)

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/posts/{id}/comments", controllers.GetComments)
	getRouter.HandleFunc("/posts/{id}", controllers.GetPost)
	getRouter.HandleFunc("/posts", controllers.GetPosts)
	getRouter.HandleFunc("/users", controllers.GetUsers)
	getRouter.HandleFunc("/users/{id}", controllers.GetUser)
	getRouter.Use(middlewares.TokenAuthMiddleware)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/users/{id}", controllers.DeleteUser)
	deleteRouter.HandleFunc("/posts/{id}/comments/{comment_id}", controllers.DeleteComment)
	deleteRouter.HandleFunc("/posts/{id}", controllers.DeletePost)
	deleteRouter.Use(middlewares.TokenAuthMiddleware)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/users/{id}", controllers.UpdateUser)
	putRouter.HandleFunc("/post/{id}", controllers.UpdatePost)
	putRouter.HandleFunc("/posts/{id}/comments/{comment_id}", controllers.UpdateComment)
	putRouter.Use(middlewares.TokenAuthMiddleware)
}
