// +build integration

package api

import "os"

var testApi = NewApi(os.Getenv("TEST_NET_CHAIN_API_HOST_URL"))
