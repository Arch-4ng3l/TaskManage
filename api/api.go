package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Arch-4ng3l/TaskManage/storage"
	"github.com/Arch-4ng3l/TaskManage/types"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listeningAddr string
	store         storage.Storage
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func NewAPIServer(addr string, store storage.Storage) *APIServer {
	return &APIServer{
		addr,
		store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", createHTTPHandleFunc(s.handleAccountRequest))
	router.HandleFunc("/login", createHTTPHandleFunc(s.handleLoginRequest))
	router.HandleFunc("/task", createHTTPHandleFunc(s.handleTaskRequest))
	router.HandleFunc("/task/{name}", createHTTPHandleFunc(s.handleGetTask))
	fmt.Println("Listening on ", s.listeningAddr)

	http.ListenAndServe(s.listeningAddr, router)
}

func (s *APIServer) handleTaskRequest(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateTask(w, r)
	case "GET":
		return s.handleGetTask(w, r)

	default:
		return fmt.Errorf("Method %s Not Allowed", r.Method)
	}
}
func (s *APIServer) handleGetTask(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleGetTaskByName(w, r)
	case "GET":
		return s.handleGetAllTasks(w, r)

	default:
		return fmt.Errorf("Method %s Not Allowed", r.Method)
	}
}

func (s *APIServer) handleCreateTask(w http.ResponseWriter, r *http.Request) error {

	req := &types.CreateTaskRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	// TODO AUTH if user that Creates the task exists and has Perms

	task := types.NewTask(req.Name, req.TaskName, req.TaskContent)
	if err := s.store.AddNewTask(task); err != nil {
		return WriteJSON(w, http.StatusConflict, NewAPIError(err.Error()))
	}
	return WriteJSON(w, http.StatusCreated, task)

}

func (s *APIServer) handleGetTaskByName(w http.ResponseWriter, r *http.Request) error {

	// tokenStr := r.Header.Get("jwt-token")
	// Auth jwt Token

	name := mux.Vars(r)["name"]

	req := &types.GetTaskRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println(err.Error())
		return err
	}

	if req.Name != name {
		return WriteJSON(w, http.StatusTeapot, NewAPIError("WTF"))
	}

	task, err := s.store.TaskFromUser(req.Name, req.TaskName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return WriteJSON(w, http.StatusOK, task)
}

func (s *APIServer) handleGetAllTasks(w http.ResponseWriter, r *http.Request) error {
	name := mux.Vars(r)["name"]

	// tokenStr := r.Header.Get("jwt-token")
	// Authenticate Token

	tasks, err := s.store.AllTasksFromUser(name)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return WriteJSON(w, http.StatusOK, tasks)
}

func (s *APIServer) handleAccountRequest(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleCreateAccount(w, r)
	default:
		WriteJSON(w, http.StatusMethodNotAllowed, fmt.Errorf("Method %s not Allowed", r.Method))
	}
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {

	req := &types.CreateAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	acc := types.NewAccount(req.Email, req.Username, req.Password)
	if err := s.store.AddNewAccount(acc); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, acc)
}

func (s *APIServer) handleDeleteRequest(w http.ResponseWriter, r *http.Request) error {
	req := &types.DeleteAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return WriteJSON(w, http.StatusBadRequest, NewAPIError("Couldn't Decode JSON"))
	}

	acc, err := s.store.GetAccountByEmail(req.Email)
	if err != nil {
		return WriteJSON(w, http.StatusForbidden, NewAPIError("Access denied"))
	}

	if acc.Password != types.CreateHash(req.Password) {
		return WriteJSON(w, http.StatusForbidden, NewAPIError("Access denied"))
	}

	if err := s.store.RemoveAccount(acc); err != nil {
		return WriteJSON(w, http.StatusNotFound, NewAPIError("Access denied"))
	}

	return WriteJSON(w, http.StatusOK, "Account Deleted")
}

func (s *APIServer) handleLoginRequest(w http.ResponseWriter, r *http.Request) error {
	req := &types.LoginAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	acc, err := s.store.GetAccountByEmail(req.Email)

	if err != nil {
		return WriteJSON(w, http.StatusForbidden, NewAPIError("Not Allowed"))
	}

	if types.CreateHash(req.Password) != acc.Password {
		return WriteJSON(w, http.StatusForbidden, NewAPIError("Not Allowed"))
	}

	return nil
}

func createHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, err)
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

func NewAPIError(s string) APIError {
	return APIError{
		s,
	}
}
