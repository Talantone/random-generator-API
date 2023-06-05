package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"random-generator-API/models"
	"random-generator-API/pkg/usecase"
	"testing"
)

func TestGenerate(t *testing.T) {
	r := gin.Default()
	//group := r.Group("/api", func(c *gin.Context) {})
	uc := new(usecase.GeneratorUseCaseMock)
	handler := NewHandler(uc)
	handler.RegisterHTTPEndpoints()

	input := models.Amount{
		Amount: 10,
	}

	body, err := json.Marshal(input)
	assert.NoError(t, err)

	uc.On("Generate", input.Amount).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/generator", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGet(t *testing.T) {

	r := gin.Default()
	//group := r.Group("/api", func(c *gin.Context) {})

	uc := new(usecase.GeneratorUseCaseMock)
	handler := NewHandler(uc)
	handler.RegisterHTTPEndpoints()
	res := models.RandomItem{
		Result: "testresult",
	}
	uc.On("Result").Return(res, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/generator", nil)
	r.ServeHTTP(w, req)

	expectedOutBody, err := json.Marshal("testresult")
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}
