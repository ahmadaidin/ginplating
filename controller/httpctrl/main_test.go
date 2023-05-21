package httpctrl

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ahmadaidin/ginplating/config"
	"github.com/ahmadaidin/ginplating/controller/httpctrl/bookctrl"
	"github.com/ahmadaidin/ginplating/domain/repository"
	"github.com/ahmadaidin/ginplating/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindAllBook(t *testing.T) {
	os.Setenv("ENV", "test")
	cfgLoader, err := config.NewLoaderAndLoad()
	require.NoError(t, err, "error when load config")

	mongoDb := database.NewMongoDatabase(cfgLoader.Config().DatabaseURI, 10)

	bookRepo := repository.NewBookRepository(mongoDb)
	ctrl := bookctrl.NewBookController(
		&cfgLoader,
		bookRepo,
	)
	handler := NewGinHttpHandler(*ctrl)

	router := handler.GetEngine()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
