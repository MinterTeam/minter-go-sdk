package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

const LatestBlockHeight = 0

type Api struct {
	client *resty.Client
}

// Create MinterAPI instance.
func NewApi(hostUrl string) *Api {
	return NewApiWithClient(hostUrl, resty.New().SetTimeout(time.Minute))
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

type ResponseError struct {
	*resty.Response
}

func NewResponseError(response *resty.Response) *ResponseError {
	return &ResponseError{Response: response}
}

func (res *ResponseError) Error() string {
	detailError := map[string]interface{}{
		"status_code": res.StatusCode(),
		"status":      res.Status(),
		"time":        res.Time(),
		"received_at": res.ReceivedAt(),
		"headers":     res.Header(),
		"body":        res,
	}
	marshal, _ := json.Marshal(detailError)
	return string(marshal)
}
