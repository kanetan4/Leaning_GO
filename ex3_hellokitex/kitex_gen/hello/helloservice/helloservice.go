// Code generated by Kitex v0.14.1. DO NOT EDIT.

package helloservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	hello "ex3_hellokitex/kitex_gen/hello"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"SayHello": kitex.NewMethodInfo(
		sayHelloHandler,
		newHelloServiceSayHelloArgs,
		newHelloServiceSayHelloResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	helloServiceServiceInfo                = NewServiceInfo()
	helloServiceServiceInfoForClient       = NewServiceInfoForClient()
	helloServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return helloServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return helloServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return helloServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "HelloService"
	handlerType := (*hello.HelloService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "hello",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.14.1",
		Extra:           extra,
	}
	return svcInfo
}

func sayHelloHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*hello.HelloServiceSayHelloArgs)
	realResult := result.(*hello.HelloServiceSayHelloResult)
	success, err := handler.(hello.HelloService).SayHello(ctx, realArg.Name)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newHelloServiceSayHelloArgs() interface{} {
	return hello.NewHelloServiceSayHelloArgs()
}

func newHelloServiceSayHelloResult() interface{} {
	return hello.NewHelloServiceSayHelloResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SayHello(ctx context.Context, name string) (r string, err error) {
	var _args hello.HelloServiceSayHelloArgs
	_args.Name = name
	var _result hello.HelloServiceSayHelloResult
	if err = p.c.Call(ctx, "SayHello", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
