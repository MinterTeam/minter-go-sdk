package api

import (
	"github.com/go-resty/resty/v2"
)

type Api struct {
	client *resty.Client
}

func NewApi(hostUrl string) *Api {
	return &Api{client: resty.New().SetHostURL(hostUrl)}
}
