package resttest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v10"
	"net/http"
	"strconv"
)

func NewServerApiMock() *gin.Engine {
	r := gin.Default()
	r.POST("/articles", mockCreateArticle)
	r.GET("/article/:id", mockGetArticle)
	return r
}

func mockCreateArticle (c *gin.Context) {
	article := &Article{}

	err := c.BindJSON(article)
	if err != nil {
		c.Data(http.StatusBadRequest, "application/json", errResponseBody(err))
		return
	}

	err = validator.New().Struct(article)
	if err != nil {
		c.Data(http.StatusBadRequest, "application/json", errResponseBody(err))
		return
	}

	article.Id = 1
	c.Header("Location", "/article/1")

	// We could build the response body by this way:
	// c.JSON(http.StatusCreated, article)
	// But it is suppose to be a copy and paste from a service API documentation
	c.Data(http.StatusCreated, "application/json",
		[]byte(`{"id":1,"author":"Author Name","title":"Title","Body":"Article content"}`))
}

func mockGetArticle (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Data(http.StatusBadRequest, "application/json", errResponseBody(err))
		return
	}

	// emulate not found
	if id == 2 {
		c.Data(http.StatusNotFound, "application/json",
			errResponseBody(errors.New("no article found by id 2")))
		return
	}

	c.Data(http.StatusOK, "application/json",
		[]byte(`{"id":1,"author":"Author Name","title":"Title","Body":"Article content"}`))

}

func errResponseBody(err error) []byte {
	return []byte(`{"error":"`+err.Error()+`"}`)
}
