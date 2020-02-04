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

func hasError(res *resty.Response, err error) error {
	if res.IsSuccess() && err == nil {
		return nil
	}

	var detailError string
	ti := res.Request.TraceInfo()
	detailError = fmt.Sprintln("Response Info:") +
		fmt.Sprintln("Error      :", err) +
		fmt.Sprintln("Status Code:", res.StatusCode()) +
		fmt.Sprintln("Status     :", res.Status()) +
		fmt.Sprintln("Time       :", res.Time()) +
		fmt.Sprintln("Received At:", res.ReceivedAt()) +
		fmt.Sprintln("Body       :\n", res) +
		fmt.Sprintln() +
		fmt.Sprintln("Request Trace Info:") +
		fmt.Sprintln("DNSLookup    :", ti.DNSLookup) +
		fmt.Sprintln("ConnTime     :", ti.ConnTime) +
		fmt.Sprintln("TLSHandshake :", ti.TLSHandshake) +
		fmt.Sprintln("ServerTime   :", ti.ServerTime) +
		fmt.Sprintln("ResponseTime :", ti.ResponseTime) +
		fmt.Sprintln("TotalTime    :", ti.TotalTime) +
		fmt.Sprintln("IsConnReused :", ti.IsConnReused) +
		fmt.Sprintln("IsConnWasIdle:", ti.IsConnWasIdle) +
		fmt.Sprintln("ConnIdleTime :", ti.ConnIdleTime)

	return errors.New(detailError)
}
