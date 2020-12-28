package viabusgo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"path"
)

const (
	apiBaseURL = "https://api.viabus.io"

	apiEndpointRegisAnon = "/auth/regisAnon"
)

type client struct {
	httpClient *http.Client
	baseURL    *url.URL
	basicUser  string
	basicPass  string
	authMode   string
}

func (c *client) Anonymous() {
	c.authMode = "basic"

	uid := uuid.New().String()
	h := hmac.New(sha256.New, []byte(uid))
	h.Write([]byte(uid))
	pass := base64.StdEncoding.EncodeToString(h.Sum(nil))

	c.basicUser = base64.StdEncoding.EncodeToString([]byte(uid))
	c.basicPass = base64.StdEncoding.EncodeToString([]byte(pass))
}

func (c client) do(req *http.Request, response interface{}) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}
func (c client) url(endpoint string) string {
	c.baseURL.Path = path.Join(c.baseURL.Path, endpoint)
	return c.baseURL.String()
}

func (c client) RegisterAnonymous() (RegisterAnonymousResponse, error) {
	postBody, _ := json.Marshal(map[string]string{
		"native_first_name":   "",
		"native_middle_name":  "",
		"native_last_name":    "",
		"english_first_name":  "",
		"english_middle_name": "",
		"english_last_name":   "",
		"gender":              "n",
		"phone":               "",
		"email":               "",
	})

	response := RegisterAnonymousResponse{}
	if c.basicUser == "" || c.basicPass == "" {
		return RegisterAnonymousResponse{}, errors.New("missing parameters")
	}

	request, err := http.NewRequest("POST", c.url(apiEndpointRegisAnon), bytes.NewBuffer(postBody))
	if err != nil {
		return response, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Basic %s:%s", c.basicUser, c.basicPass))

	err = c.do(request, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

type Client interface {
	RegisterAnonymous() (RegisterAnonymousResponse, error)
	Anonymous()
}

type ClientOption func(c *client) error

func New(options ...ClientOption) (Client, error) {
	c := &client{
		httpClient: http.DefaultClient,
	}

	u, err := url.ParseRequestURI(apiBaseURL)
	if err != nil {
		return nil, err
	}
	c.baseURL = u

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
