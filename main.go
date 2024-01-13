package main

import (
	"templ_test/backend"
)

func main() {
	backend.ConnectDB()
	backend.SetUpHandlers()
}
