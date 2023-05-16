package httpctrl

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmadaidin/ginplating/controller/httpctrl/bookctrl"
	"github.com/ahmadaidin/ginplating/domain/repository"
	"github.com/ahmadaidin/ginplating/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestFindAllBook(t *testing.T) {
	mongoDb := database.NewMongoDatabase("mongodb://mongodb:27017/ginplating", 10)

	bookRepo := repository.NewBookRepository(mongoDb)
	ctrl := bookctrl.NewBookController(
		*bookRepo,
	)
	handler := NewGinHttpHandler(*ctrl)

	router := handler.GetEngine()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
