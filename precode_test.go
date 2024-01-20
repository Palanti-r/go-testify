package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 5
	city := "moscow"
	req, err := http.NewRequest(http.MethodGet, "/cafe?count="+strconv.Itoa(totalCount)+"&city="+city, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
	// здесь нужно добавить необходимые проверки
}
func TestMainHandlerWhenRequestIsCorrectly(t *testing.T) {
	totalCount := 4
	city := "moscow"
	req, err := http.NewRequest(http.MethodGet, "/cafe?count="+strconv.Itoa(totalCount)+"&city="+city, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEqual(t, 0, responseRecorder.Result().ContentLength)

	// здесь нужно добавить необходимые проверки
}
func TestMainHandlerWhenCityIsUnSupported(t *testing.T) {
	totalCount := 4
	city := "Sochi"
	req, err := http.NewRequest(http.MethodGet, "/cafe?count="+strconv.Itoa(totalCount)+"&city="+city, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Contains(t, responseRecorder.Body.String(), "wrong city value")
	// здесь нужно добавить необходимые проверки
}
