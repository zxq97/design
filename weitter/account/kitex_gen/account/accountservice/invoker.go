// Code generated by Kitex v0.5.0. DO NOT EDIT.

package accountservice

import (
	server "github.com/cloudwego/kitex/server"
	account "github.com/zxq97/design/weitter/account/kitex_gen/account"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler account.AccountService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
