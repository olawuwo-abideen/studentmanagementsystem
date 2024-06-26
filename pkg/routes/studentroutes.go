package routes

import (
	"github.com/gorilla/mux"
	"github.com/olawuwo-abideen/studentmanagementsystem/pkg/controllers"
)

var StudentManagementRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/students/", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")
}
