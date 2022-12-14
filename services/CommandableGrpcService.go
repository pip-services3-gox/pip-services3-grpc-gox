package services

import (
	"context"

	ccomands "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
)

// CommandableGrpcService abstract service that receives commands via GRPC protocol
// to operations automatically generated for commands defined in ICommandable components.
// Each command is exposed as invoke method that receives command name and parameters.
//
// Commandable services require only 3 lines of code to implement a robust external
// GRPC-based remote interface.
//
//	Configuration parameters:
//
//		- dependencies:
//			- endpoint:              override for HTTP Endpoint dependency
//			- controller:            override for Controller dependency
//		- connection(s):
//			- discovery_key:         (optional) a key to retrieve the connection from  IDiscovery
//			- protocol:              connection protocol: http or https
//			- host:                  host name or IP address
//			- port:                  port number
//			- uri:                   resource URI or connection string with all parameters in it
//
//	References:
//
//		- *:logger:*:*:1.0               (optional) ILogger components to pass log messages
//		- *:counters:*:*:1.0             (optional) ICounters components to pass collected measurements
//		- *:discovery:*:*:1.0            (optional) IDiscovery services to resolve connection
//		- *:endpoint:grpc:*:1.0          (optional) GrpcEndpoint reference
//
// See CommandableGrpcClient
// See GrpcService
//
// Example:

//	type MyCommandableGrpcService struct {
//		*CommandableGrpcService
//	}
//	func NewCommandableGrpcService() *CommandableGrpcService {
//		c := DummyCommandableGrpcService{}
//		c.CommandableGrpcService = grpcservices.NewCommandableGrpcService("myservice")
//		c.DependencyResolver.Put("controller", cref.NewDescriptor("mygroup", "controller", "default", "*", "*"))
//		return &c
//	}
//
//	service := NewMyCommandableGrpcService();
//	service.Configure(ctx, cconf.NewConfigParamsFromTuples(
//		"connection.protocol", "http",
//		"connection.host", "localhost",
//		"connection.port", "8080",
//	));
//	service.SetReferences(ctx, cref.NewReferencesFromTuples(
//		cref.NewDescriptor("mygroup","controller","default","default","1.0"), controller
//	));
//
//	opnErr := service.Open(ctx, "123")
//	if opnErr == nil {
//		fmt.Println("The GRPC service is running on port 8080");
//	}
//
type CommandableGrpcService struct {
	*GrpcService
	name       string
	commandSet *ccomands.CommandSet
}

// NewCommandableGrpcService method are creates a new instance of the service.
//    - name a service name.
func InheritCommandableGrpcService(overrides IGrpcServiceOverrides, name string) *CommandableGrpcService {
	c := &CommandableGrpcService{}
	c.GrpcService = InheritGrpcService(overrides, "")
	c.name = name
	c.DependencyResolver.Put(context.Background(), "controller", "none")
	return c
}

// Register method are registers all service command in gRPC endpoint.
func (c *CommandableGrpcService) Register() {

	resCtrl, depErr := c.DependencyResolver.GetOneRequired("controller")
	if depErr != nil {
		return
	}
	controller, ok := resCtrl.(ccomands.ICommandable)
	if !ok {
		c.Logger.Error(context.Background(), "CommandableHttpService", nil, "Can't cast Controller to ICommandable")
		return
	}
	c.commandSet = controller.GetCommandSet()

	commands := c.commandSet.Commands()
	var index = 0
	for index = 0; index < len(commands); index++ {
		command := commands[index]

		method := c.name + "." + command.Name()

		c.RegisterCommandableMethod(method, nil,
			func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
				timing := c.Instrument(ctx, correlationId, method)
				res, err := command.Execute(ctx, correlationId, args)
				timing.EndTiming(ctx, err)
				return res, err
			})
	}
}
