package resttest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ApiClient struct {
	BaseUrl string
}

func (c *ApiClient) CreateArticle(a *Article) (int, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", c.BaseUrl+"/articles", bytes.NewReader(b))
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusCreated {
		return 0, errors.New("not expected status code")
	}

	buf := &bytes.Buffer{}
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return 0, err
	}

	article := &Article{}
	err = json.Unmarshal(buf.Bytes(), article)

	return article.Id, err
}
