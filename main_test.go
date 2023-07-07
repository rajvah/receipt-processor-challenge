package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateReciept(t *testing.T) {
	r := SetUpRouter()

	r.POST("/receipts/process", CreateReciept())

	id := uuid.New().String()
	reciept := Receipt{
		ID:           id,
		Retailer:     "Test",
		PurchaseDate: "2022-02-01",
		PurchaseTime: "11:45",
		Items:        []Item{},
		Total:        "23.06",
	}
	jsonValue, _ := json.Marshal(reciept)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateRecieptWithIncorrectDateFormat(t *testing.T) {
	r := SetUpRouter()

	r.POST("/receipts/process", CreateReciept())

	id := uuid.New().String()
	reciept := Receipt{
		ID:           id,
		Retailer:     "Test",
		PurchaseDate: "2022-20-01",
		PurchaseTime: "11:45",
		Items:        []Item{},
		Total:        "23.06",
	}
	jsonValue, _ := json.Marshal(reciept)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateRecieptWithIncorrectTimeFormat(t *testing.T) {
	r := SetUpRouter()

	r.POST("/receipts/process", CreateReciept())

	id := uuid.New().String()
	reciept := Receipt{
		ID:           id,
		Retailer:     "Test",
		PurchaseDate: "2022-02-01",
		PurchaseTime: "31:45",
		Items:        []Item{},
		Total:        "23.06",
	}
	jsonValue, _ := json.Marshal(reciept)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

//Todo:
// Gin Context is not setting the parameters correctly, get response TEST returns 404. Only happens in the test suite

//func TestGetPoints(t *testing.T) {
//	r := SetUpRouter()
//
//	r.POST("/receipts/process", CreateReciept())
//
//	//id := uuid.New().String()
//	reciept := Receipt{
//		Retailer:     "Test",
//		PurchaseDate: "2022-02-01",
//		PurchaseTime: "11:45",
//		Items:        []Item{},
//		Total:        "23.06",
//	}
//	jsonValue, _ := json.Marshal(reciept)
//	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))
//
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	responseData, _ := ioutil.ReadAll(w.Body)
//	assert.Equal(t, http.StatusCreated, w.Code)
//
//	json.Unmarshal(responseData, &reciept)
//
//	mockResponse := `{"points": 10}`
//
//	r.GET("/receipts/:id/points", GetPoints())
//
//	URL := fmt.Sprintf("/receipts/%s/points", reciept.ID)
//	req, _ = http.NewRequest("GET", URL, nil)
//	w_get := httptest.NewRecorder()
//	r.ServeHTTP(w_get, req)
//
//	responseData, _ = ioutil.ReadAll(w_get.Body)
//	assert.Equal(t, mockResponse, string(responseData))
//	assert.Equal(t, http.StatusNotFound, w_get.Code)
//}

func TestGetPointsIDNotFound(t *testing.T) {
	r := SetUpRouter()

	r.GET("/receipts/:id/points", GetPoints())

	URL := fmt.Sprintf("/receipts/%s/points", uuid.New().String())
	req, _ := http.NewRequest("GET", URL, nil)
	w_get := httptest.NewRecorder()
	r.ServeHTTP(w_get, req)

	assert.Equal(t, http.StatusNotFound, w_get.Code)

}
