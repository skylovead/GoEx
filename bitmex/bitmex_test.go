package bitmex

import (
	"github.com/nntaoli-project/GoEx"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "socks5",
				Host:   "127.0.0.1:1080"}, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}
var bm = New(httpProxyClient, "", "")

func TestBitmex_GetDepth(t *testing.T) {
	return
	dep, err := bm.GetDepth(2, goex.NewCurrencyPair(goex.XBT, goex.USD))
	assert.Nil(t, err)
	t.Log(dep)
}

func TestBitmex_GetAccount(t *testing.T) {
	acc, err := bm.GetAccount()
	t.Log(acc, err)
}
