package test

import (
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type TestGrpcClient struct {
	*clients.GrpcClient
}

func NewTestGrpcClient(name string) *TestGrpcClient {
	c := &TestGrpcClient{}
	c.GrpcClient = clients.NewGrpcClient(name)
	return c
}
