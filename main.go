package main

import (
	"fmt"
	"net/http"
	"rest-api-simple-auth/controllers"
	"rest-api-simple-auth/core"
	"rest-api-simple-auth/middleware"
)

func init() {
	fmt.Println("Init func from package main")
}
func main() {

	http.Handle("/ping", middleware.Auth(http.HandlerFunc(controllers.Pong)))

	app := core.App{}
	app.Setup()
	app.ListenAndServe()
}
