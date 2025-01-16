// Code generated by Kitex v0.9.1. DO NOT EDIT.
package shortlinkservice

import (
	shortlink "github.com/FinnTew/FinnEcommerce/src/internal/client/kitex_gen/shortlink"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler shortlink.ShortLinkService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler shortlink.ShortLinkService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
