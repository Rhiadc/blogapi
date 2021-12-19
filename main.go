package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rhiadc/blogapi/api"
	"github.com/rhiadc/blogapi/database/redis"
	"github.com/rhiadc/blogapi/routes"
)

func main() {

	api.Run()

	sm := mux.NewRouter()
	routes.InitRoutes(sm)

	fmt.Println("Server running on port 8080")
	redis.InitRedis()
	http.ListenAndServe(":8080", sm)

}
