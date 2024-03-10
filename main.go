package main

import (
	"templ_test/internal/persistance"
	backend "templ_test/internal/server"
)

func main() {
	db := persistance.ConnectDB()
	app := backend.NewApp(db)
	app.SetUpHandlers()
	app.Start()
}
