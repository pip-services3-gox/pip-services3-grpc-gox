package test

import (
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type TestCommandableGrpcClient struct {
	clients.CommandableGrpcClient
}

func NewTestCommandableGrpcClient(name string) *TestCommandableGrpcClient {
	c := &TestCommandableGrpcClient{}
	c.CommandableGrpcClient = *clients.NewCommandableGrpcClient(name)
	return c
}
