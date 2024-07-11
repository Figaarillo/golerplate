package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Figaarillo/golerplate/internal/shared/config"
	"github.com/Figaarillo/golerplate/internal/shared/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	Router *mux.Router
	DB     *gorm.DB
)

func InitDBAndRouter(t *testing.T) {
	env, err := config.NewEnvConf("../../.env.test")
	if err != nil {
		t.Fatalf("Error: Cannot read .env.test file: %v", err)
	}

	if DB == nil {
		DB = config.InitGorm(env)
	}

	if Router == nil {
		Router = config.InitRouter()
	}
}

func ClearAllTables(t *testing.T, tables []string) {
	for _, table := range tables {
		if err := DB.Exec("DELETE FROM " + table).Error; err != nil {
			t.Fatalf("Could not clear table %s: %v", table, err)
		}
	}
}

func NewHTTPRequest(t *testing.T, method, url string, body interface{}) *http.Request {
	var reqBody *bytes.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Error: Cannot marshal request body: %v", err)
		}
		reqBody = bytes.NewReader(jsonBody)
	} else {
		reqBody = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		t.Fatalf("Error: Cannot create new request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func ExecuteHTTPRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Router.ServeHTTP(rr, req)
	return rr
}

func AssertResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Error: Expected response code %d. Got %d\n", expected, actual)
	}
}

func AssertEqual(t *testing.T, name string, expected, got interface{}) {
	if got != expected {
		t.Errorf("Error: Expected %s to be '%v', got '%v'", name, expected, got)
	}
}

func CreateNewEntity[T any](t *testing.T, url string, entity T) {
	req := NewHTTPRequest(t, http.MethodPost, url, entity)
	res := ExecuteHTTPRequest(req)
	AssertResponseCode(t, http.StatusCreated, res.Code)
}

func UnmarshalResponse(t *testing.T, res *httptest.ResponseRecorder) utils.Response {
	var response utils.Response
	if err := json.Unmarshal(res.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error: Cannot unmarshal response body: %v", err)
	}

	return response
}

func ParseResponseData[T any](t *testing.T, res *httptest.ResponseRecorder) []T {
	response := UnmarshalResponse(t, res)

	var entityList []T
	switch body := response.Body.(type) {
	case map[string]interface{}:
		entityJSON, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Error: Cannot marshal entity JSON: %v", err)
		}

		var entity T
		if err := json.Unmarshal(entityJSON, &entity); err != nil {
			t.Fatalf("Error: Cannot unmarshal entity JSON: %v", err)
		}

		entityList = append(entityList, entity)
	case []interface{}:
		for _, e := range body {
			entityJSON, err := json.Marshal(e)
			if err != nil {
				t.Fatalf("Error: Cannot marshal entity JSON: %v", err)
			}

			var entity T
			if err := json.Unmarshal(entityJSON, &entity); err != nil {
				t.Fatalf("Error: Cannot unmarshal entity JSON: %v", err)
			}

			entityList = append(entityList, entity)
		}
	default:
		t.Fatalf("Error: Unexpected response body type: %T", response.Body)
	}

	return entityList
}
