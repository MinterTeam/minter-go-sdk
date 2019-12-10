package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const LatestBlockHeight = 0

type Api struct {
	client *resty.Client
}

// Create MinterAPI instance.
func NewApi(hostUrl string) *Api {
	return &Api{client: resty.New().SetHostURL(hostUrl)}
}

// Create MinterAPI instance with custom client
func NewApiWithClient(hostUrl string, client *resty.Client) *Api {
	return &Api{client: client.SetHostURL(hostUrl)}
}

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: \"%s\", data: \"%s\"", e.Code, e.Message, e.Data)
}
