package resttest

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	baseUrl = "http://localhost:8123"
)

func TestMain(m *testing.M) {
	go func() {
		err := NewServerApiMock().Run(":8123")
		if err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}

func TestClientCanPostArticle(t *testing.T) {
	cli := &ApiClient{BaseUrl: baseUrl}

	article := &Article{
		Title:  "Title",
		Author: "Author Name",
		Body:   "Article content",
	}

	id, err := cli.CreateArticle(article)
	assert.Nil(t, err)
	assert.Equal(t, 1, id)
}

func TestClientCanHandleErrorWhilePostingArticle(t *testing.T) {
	cli := &ApiClient{BaseUrl: baseUrl}

	article := &Article{}

	id, err := cli.CreateArticle(article)
	assert.NotNil(t, err)
	assert.Equal(t, "not expected status code", err.Error())
	assert.Equal(t, 0, id)
}
