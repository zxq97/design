// Code generated by Kitex v0.5.0. DO NOT EDIT.
package likeservice

import (
	server "github.com/cloudwego/kitex/server"
	like "github.com/zxq97/design/weitter/like/kitex_gen/like"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler like.LikeService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
