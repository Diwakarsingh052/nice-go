package data

import "net/http"

type Client interface {
	DoReq(req *http.Request) (*http.Response, error)
}
