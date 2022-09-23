package test_services

import (
	"context"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	grpcservices "github.com/pip-services3-gox/pip-services3-grpc-gox/services"
)

type DummyCommandableGrpcService struct {
	*grpcservices.CommandableGrpcService
}

func NewDummyCommandableGrpcService() *DummyCommandableGrpcService {
	c := &DummyCommandableGrpcService{}
	c.CommandableGrpcService = grpcservices.InheritCommandableGrpcService(c, "dummy")
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("pip-services-dummies", "controller", "default", "*", "*"))
	return c
}
