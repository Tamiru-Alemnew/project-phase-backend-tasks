package main

import (
	"github.com/Tamiru-Alemnew/project-phase-backend-tasks/Task-three/library_management/controllers"
	"github.com/Tamiru-Alemnew/project-phase-backend-tasks/Task-three/library_management/services"
)

func main() {

    library := services.NewLibrary()
    controller := controllers.NewLibraryController(library)
    controller.RunLibrarySystem()
}

