package gateway

import (
	"net/http/httputil"
	"net/url"
)

type Gateway struct {
	hosts map[string]*httputil.ReverseProxy
}

func NewGateway() *Gateway {
	return &Gateway{
		hosts: make(map[string]*httputil.ReverseProxy),
	}
}

func (g *Gateway) AddHost(t, path string) error {
	target, err := url.Parse(t)
	if err != nil {
		return err
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	g.hosts[path] = proxy

	return nil
}
