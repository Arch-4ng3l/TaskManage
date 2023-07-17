package www

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Arch-4ng3l/TaskManage/types"
	"github.com/gorilla/mux"
)

func AddFrontend(route *mux.Router) {

	route.HandleFunc("/", createFileHandleFunc("index.html"))
	route.HandleFunc("/dashboard", handleDashboard)
	route.HandleFunc("/login", createFileHandleFunc("login.html"))
	route.HandleFunc("/signup", createFileHandleFunc("signup.html"))

	fs := http.FileServer(http.Dir("./www/static"))

	route.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

}

func createFileHandleFunc(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := "./www/static/" + filename
		http.ServeFile(w, r, path)
	}
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./www/static/dashboard.html"))
	//Validate Account
	acc := types.NewAccount("Test", "Test", "Test")
	if err := tmpl.Execute(w, *acc); err != nil {
		fmt.Println(err.Error())
	}

}
