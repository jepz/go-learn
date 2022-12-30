package main

import (
	"net/http"

	"github.com/jepz/learn-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) //nil means default in this case
}
