package bookctrl

import (
	"github.com/ahmadaidin/ginplating/domain/dto"
	"github.com/ahmadaidin/ginplating/domain/repository"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookRepo repository.BookRepository
}

func NewBookController(
	bookRepo repository.BookRepository,
) *BookController {
	return &BookController{
		bookRepo: bookRepo,
	}
}

// @Summary Find all books
// @Description Find all books
// @Tags Book
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.Book
// @Router /books [get]
func (ctr *BookController) FindAll(c *gin.Context) {
	ctx := c.Request.Context()
	opt := dto.FindAllBookOptions{}

	c.BindQuery(&opt)
	books, _, err := ctr.bookRepo.FindAll(ctx, opt)
	if err.Valid() {
		err.PrependMsg("error in BookController.FindAll when calling ctr.bookRepo.FindAll")
		c.AbortWithError(err.HttpStatus(), err.Err())
	}

	c.JSON(200, books)
}
