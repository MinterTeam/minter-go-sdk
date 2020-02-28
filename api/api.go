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
	detailError := map[string]string{
		"status_code": fmt.Sprintf("%d", res.StatusCode()),
		"status":      res.Status(),
		"time":        fmt.Sprintf("%f seconds", res.Time().Seconds()),
		"received_at": fmt.Sprintf("%v", res.ReceivedAt()),
		"headers":     fmt.Sprintf("%v", res.Header()),
		"body":        res.String(),
	}
	marshal, _ := json.Marshal(detailError)

	return string(marshal)
}
