package user

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

func TestGetUserRoles(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/users/alice/roles", nil)
	if err != nil {
		t.Fatal("Request failed")
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.Handle("/v1/users/{user}/roles", negroni.New(negroni.HandlerFunc(GetUserRoles))).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d \n. Got %d", http.StatusOK, status)
	}
}

func TestDeleteRoleFromUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/users/alice/roles/visitante3", nil)
	if err != nil {
		t.Fatal("Request failed")
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.Handle("/v1/users/{user}/roles/{role}", negroni.New(negroni.HandlerFunc(DeleteRoleFromUser))).Methods("DELETE")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d \n. Got %d", http.StatusOK, status)
	}
}

func TestCheckUserPermission(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/users/alice/resource/data1/permission/read", nil)
	if err != nil {
		t.Fatal("Request failed")
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.Handle("/v1/users/{subject}/resource/{object}/permission/{action}", negroni.New(negroni.HandlerFunc(CheckUserPermission))).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d \n. Got %d", http.StatusOK, status)
	}
}