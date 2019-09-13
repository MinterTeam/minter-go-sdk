package api

import (
	"strings"
)

type Api struct {
	hostUrl string
}

func NewApi(hostUrl string) *Api {
	return &Api{hostUrl: strings.TrimRight(hostUrl, "/")}
}
