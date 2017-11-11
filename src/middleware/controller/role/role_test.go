package role

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	"strings"
	"middleware/service/user"
	"middleware/service/role"
)

func TestGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/roles", nil)
	if err != nil {
		t.Fatal("Request failed")
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.Handle("/v1/roles", negroni.New(negroni.HandlerFunc(GetAll))).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d \n. Got %d", http.StatusOK, status)
	}

	assert.NotEqual(t, -1, strings.Index(string(rr.Body.String()), "roles"), "Response body differs")
}

func TestCreate(t *testing.T) {

	user.DeleteRoleFromUser("alice", "visitante")
	router := mux.NewRouter()
	router.Handle("/v1/roles", negroni.New(negroni.HandlerFunc(Create))).Methods("POST")

	dataSend := []string{
		`{"user":"alice", "role":"visitante"}`,
		`{"user":"alice", "role":"visitante"}`,
		`{"user":"alice", "role":""}`,
		`{"user":"", "role":"visitante"}`,
		`{"user":"", "role":""}`,
		`{"user":"", "role":""`,
	}

	dataResponse := map[int]string{
		0: "Role criado com sucesso para o usuário : alice",
		1: "Role já existe para esse usuário",
		2: "Necessário usuário e role",
		3: "Necessário usuário e role",
		4: "Necessário usuário e role",
		5: "Um erro inesperado ocorreu, verifique o formato do seu json",
	}

	codeExpected := map[int]int {
		0: http.StatusOK,
		1: http.StatusOK,
		2: http.StatusBadRequest,
		3: http.StatusBadRequest,
		4: http.StatusBadRequest,
		5: http.StatusInternalServerError,
	}

	for iter := range dataSend {
		reader := strings.NewReader(dataSend[iter])

		req, err := http.NewRequest("POST", "/v1/roles", reader)
		if err != nil {
			t.Fatal("Request failed")
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != codeExpected[iter] {
			t.Errorf("Status code differs. Expected %d \n. Got %d", codeExpected[iter], status)
		}
		assert.NotEqual(t, -1, strings.Index(string(rr.Body.String()),
			dataResponse[iter]),
			"Response body differs")
	}
	user.DeleteRoleFromUser("alice", "visitante")
}

func TestGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/roles/alice", nil)
	if err != nil {
		t.Fatal("Request failed")
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.Handle("/v1/roles/{name}", negroni.New(negroni.HandlerFunc(Get))).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d \n. Got %d", http.StatusOK, status)
	}
}

func TestUpdate(t *testing.T) {

	role.DeletePermissionForUser("alice", "read")
	router := mux.NewRouter()
	router.Handle("/v1/roles", negroni.New(negroni.HandlerFunc(Update))).Methods("PUT")

	dataSend := []string{
		`{"user":"alice", "permission": "read"}`,
		`{"user":"alice", "permission": "read"}`,
		`{"user":"alice", "permission": ""}`,
		`{"user":"", "permission": "read"}`,
		`{"user":"", "permission": ""}`,
		`{"user":"alice", "permission": "read"`,
	}

	dataResponse := map[int]string{
		0: "Atualização efetuado com sucesso",
		1: "Usuário já possui essa permissão",
		2: "Necessário usuário e permissão",
		3: "Necessário usuário e permissão",
		4: "Necessário usuário e permissão",
		5: "Um erro inesperado ocorreu, verifique o formato do seu json",
	}

	codeExpected := map[int]int {
		0: http.StatusOK,
		1: http.StatusOK,
		2: http.StatusBadRequest,
		3: http.StatusBadRequest,
		4: http.StatusBadRequest,
		5: http.StatusInternalServerError,
	}

	for iter := range dataSend {
		reader := strings.NewReader(dataSend[iter])

		req, err := http.NewRequest("PUT", "/v1/roles", reader)
		if err != nil {
			t.Fatal("Request failed")
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != codeExpected[iter] {
			t.Errorf("Status code differs. Expected %d \n. Got %d", codeExpected[iter], status)
		}
		assert.NotEqual(t, -1, strings.Index(string(rr.Body.String()),
			dataResponse[iter]),
			"Response body differs")
	}
	role.DeletePermissionForUser("alice", "read")
}

func TestDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/v1/roles/administrador", nil)
	if err != nil {
		t.Fatal("Request failed")
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.Handle("/v1/roles/{name}", negroni.New(negroni.HandlerFunc(Delete))).Methods("DELETE")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d \n. Got %d", http.StatusOK, status)
	}
}