package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 5
	city := "moscow"
	req, err := http.NewRequest(http.MethodGet, "/cafe", nil)

	query := req.URL.Query()
	query.Add("city", city)
	query.Add("count", strconv.Itoa(totalCount))
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
}
func TestMainHandlerWhenRequestIsCorrectly(t *testing.T) {
	totalCount := 4
	city := "moscow"
	req, err := http.NewRequest(http.MethodGet, "/cafe", nil)

	query := req.URL.Query()
	query.Add("city", city)
	query.Add("count", strconv.Itoa(totalCount))
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEqual(t, 0, responseRecorder.Result().ContentLength)
}
func TestMainHandlerWhenCityIsUnSupported(t *testing.T) {
	totalCount := 4
	city := "UnSupported city"
	req, err := http.NewRequest(http.MethodGet, "/cafe", nil)

	query := req.URL.Query()
	query.Add("city", city)
	query.Add("count", strconv.Itoa(totalCount))
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Contains(t, responseRecorder.Body.String(), "wrong city value")
}
