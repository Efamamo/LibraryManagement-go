package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {
	var library = services.NewLibrary()
	controllers.StartConsole(library)
}
