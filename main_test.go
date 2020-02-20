package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarathep/hellogoapi/api"
	"github.com/tarathep/hellogoapi/models"
	"github.com/tarathep/hellogoapi/route"
)

type mockDB struct{}
type mockKafka struct{}

func (kafka *mockKafka) Publish(topic string, message string) error {
	return nil
}

func (db *mockDB) AllHello() ([]*models.Hello, error) {
	hellos := make([]*models.Hello, 0)
	hellos = append(hellos, &models.Hello{"C++", "c is height lv programing"})
	hellos = append(hellos, &models.Hello{"VB", "c is basic programing"})

	return hellos, nil
}

func (db *mockDB) InsertHello(hello models.Hello) (models.Hello, error) {
	return hello, nil
}

/*
func TestHelloFindAll(t *testing.T) {
	route := route.Route{api.HelloHandler{DB: &mockDB{}}}
	r := route.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	r.ServeHTTP(w, req)

	//test http status ok = 200
	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), `[{"language":"C++","message":"c is height lv programing"},{"language":"VB","message":"c is basic programing"}]`)

}
*/

func TestHelloAdd(t *testing.T) {
	route := route.Route{api.HelloHandler{DB: &mockDB{}, Kafka: &mockKafka{}}}
	r := route.SetupRouter()

	w := httptest.NewRecorder()
	inputJSON := `{"language":"C++","message":"c is height lv programing"}`
	req, _ := http.NewRequest("POST", "/hello", strings.NewReader(inputJSON))
	r.ServeHTTP(w, req)
	//test http status ok = 200
	assert.Equal(t, http.StatusOK, w.Code)

	//test body is not null
	assert.NotNil(t, w.Body)

	//test body response
	assert.Equal(t, strings.Trim(w.Body.String(), "\n"), "success")

}
