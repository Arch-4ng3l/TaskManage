package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listeningAddr string
	store         Storage
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func NewAPIServer(addr string, store Storage) *APIServer {
	return &APIServer{
		addr,
		store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", createHTTPHandleFunc(s.handleAccountRequest))
	router.HandleFunc("/login", createHTTPHandleFunc(s.handleLoginRequest))

	fmt.Println("Listening on ", s.listeningAddr)

	http.ListenAndServe(s.listeningAddr, router)
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

	req := &CreateAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	acc := NewAccount(req.Email, req.Username, req.Password)
	if err := s.store.AddNewAccount(acc); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, acc)
}

func (s *APIServer) handleDeleteRequest(w http.ResponseWriter, r *http.Request) error {
	req := &DeleteAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return WriteJSON(w, http.StatusBadRequest, "Couldnt Decode JSON")
	}

	acc, err := s.store.GetAccountByEmail(req.Email)
	if err != nil {
		return WriteJSON(w, http.StatusForbidden, "Access denied")
	}

	if acc.Password != CreateHash(req.Password) {
		return WriteJSON(w, http.StatusForbidden, "Access denied")
	}

	if err := s.store.RemoveAccount(acc); err != nil {
		return WriteJSON(w, http.StatusNotFound, "Access denied")
	}

	return WriteJSON(w, http.StatusOK, "Account Deleted")
}

func (s *APIServer) handleLoginRequest(w http.ResponseWriter, r *http.Request) error {

	req := &LoginAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	acc, err := s.store.GetAccountByEmail(req.Email)

	if err != nil {
		return WriteJSON(w, http.StatusForbidden, "Not Allowed")
	}

	if CreateHash(req.Password) != acc.Password {
		return WriteJSON(w, http.StatusForbidden, "Not Allowed")
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
