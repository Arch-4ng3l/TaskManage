package www

import (
	"html/template"
	"net/http"

	"github.com/Arch-4ng3l/TaskManage/storage"
	"github.com/Arch-4ng3l/TaskManage/types"
	"github.com/Arch-4ng3l/TaskManage/util"
	"github.com/gorilla/mux"
)

var store storage.Storage

type Data struct {
	User  *types.Account
	Tasks []*types.Task
}

func AddFrontend(s storage.Storage, route *mux.Router) {
	store = s
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

	cookies, err := util.GetCookies(r)
	if err != nil {
		return
	}
	name := cookies[0]
	email := cookies[1]
	token := cookies[2]
	acc := types.NewAccount(email, name, "")
	if !util.AuthJWT(token, acc) {
		return
	}

	tasks, _ := store.AllTasksFromUser(name)

	task := types.NewTask("TEST", "TEST", "TEST")
	tasks = append(tasks, task)

	user := &Data{
		acc,
		tasks,
	}

	if err := tmpl.Execute(w, user); err != nil {
	}

}
