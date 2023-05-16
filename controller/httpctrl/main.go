package httpctrl

import (
	"github.com/gin-gonic/gin"

	"github.com/ahmadaidin/ginplating/controller/httpctrl/bookctrl"
)

type GinHandler struct {
	Engine   *gin.Engine
	bookCtrl bookctrl.BookController
}

func (handler *GinHandler) GetEngine() *gin.Engine {
	return handler.Engine
}

func (handler *GinHandler) route() {
	bookRouter := handler.Engine.Group("books")
	bookRouter.GET("", handler.bookCtrl.FindAll)
}

func NewGinHttpHandler(
	bookctrlCtrl bookctrl.BookController,
) GinHandler {
	engine := gin.Default()
	gh := GinHandler{
		Engine:   engine,
		bookCtrl: bookctrlCtrl,
	}
	gh.route()
	return gh
}
