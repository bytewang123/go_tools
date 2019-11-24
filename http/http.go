package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func (c *Client) Init() {
	c.client = &http.Client{
		Timeout: 60 * time.Second,
	}
}

func (c *Client) PostParams(url string, params map[string]string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := c.client.Do(request)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%+v", resp.Status)
	}
	rbody := &bytes.Buffer{}
	_, err = rbody.ReadFrom(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	return rbody.String(), nil
}

func (c *Client) PostJson(url string, jsonByte []byte) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%+v", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
