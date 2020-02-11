package api

import (
	"errors"
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

func hasError(res *resty.Response) error {
	if res.IsSuccess() {
		return nil
	}

	var detailError string
	detailError = fmt.Sprintln("Response Info:") +
		fmt.Sprintln("Status Code:", res.StatusCode()) +
		fmt.Sprintln("Status     :", res.Status()) +
		fmt.Sprintln("Time       :", res.Time()) +
		fmt.Sprintln("Received At:", res.ReceivedAt()) +
		fmt.Sprintln("Headers    :", res.Header()) +
		fmt.Sprintln("Body       :\n", res)

	return errors.New(detailError)
}
