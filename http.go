package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

const simpleAuthSessKey = "_simpleauth_sess"

// Client is a wrapper around a http client
type Client struct {
	simpleAuthSess string
}

// NewClient returns a new client with the session cookie value
func NewClient(simpleAuthSess string) Client {
	return Client{simpleAuthSess}
}

// Do performs the HTTP request and returns the response
func (c Client) Do(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:     simpleAuthSessKey,
		Value:    c.simpleAuthSess,
		Path:     "/",
		Domain:   ".humblebundle.com",
		Secure:   true,
		HttpOnly: true,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
