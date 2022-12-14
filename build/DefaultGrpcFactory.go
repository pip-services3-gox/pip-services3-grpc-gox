package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	grpcservices "github.com/pip-services3-gox/pip-services3-grpc-gox/services"
)

// DefaultGrpcFactory creates GRPC components by their descriptors.
// See Factory
// See GrpcEndpoint
// See HeartbeatGrpcService
// See StatusGrpcService
type DefaultGrpcFactory struct {
	*cbuild.Factory
	Descriptor             *cref.Descriptor
	GrpcEndpointDescriptor *cref.Descriptor
	//  StatusServiceDescriptor *cref.Descriptor = new Descriptor("pip-services", "status-service", "grpc", "*", "1.0");
	//  HeartbeatServiceDescriptor *cref.Descriptor = new Descriptor("pip-services", "heartbeat-service", "grpc", "*", "1.0");
}

// NewDefaultGrpcFactory method are creates a new instance of the factory.
func NewDefaultGrpcFactory() *DefaultGrpcFactory {

	c := DefaultGrpcFactory{
		Factory: cbuild.NewFactory(),
	}
	c.Descriptor = cref.NewDescriptor("pip-services", "factory", "grpc", "default", "1.0")
	c.GrpcEndpointDescriptor = cref.NewDescriptor("pip-services", "endpoint", "grpc", "*", "1.0")
	//  c.StatusServiceDescriptor  = cref.NewDescriptor("pip-services", "status-service", "grpc", "*", "1.0");
	//  c.HeartbeatServiceDescriptor = cref.NewDescriptor("pip-services", "heartbeat-service", "grpc", "*", "1.0");

	c.RegisterType(c.GrpcEndpointDescriptor, grpcservices.NewGrpcEndpoint)
	// c.RegisterType(c.HeartbeatServiceDescriptor, grpcservices.NewHeartbeatGrpcService);
	// c.RegisterType(c.StatusServiceDescriptor, grpcservices.NewStatusGrpcService);
	return &c
}
