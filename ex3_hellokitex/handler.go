package main

import (
	"context"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// SayHello implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) SayHello(ctx context.Context, name string) (resp string, err error) {
	// TODO: Your code here...
	return "Hello " + name, nil
}
