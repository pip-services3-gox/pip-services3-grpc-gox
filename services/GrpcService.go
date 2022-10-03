package services

import (
	"context"
	"encoding/json"
	"strings"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cerr "github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	ccount "github.com/pip-services3-gox/pip-services3-components-gox/count"
	clog "github.com/pip-services3-gox/pip-services3-components-gox/log"
	ctrace "github.com/pip-services3-gox/pip-services3-components-gox/trace"
	rpcserv "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
	"google.golang.org/grpc"
)

type IGrpcServiceOverrides interface {
	Register()
}

// GrpcService abstract service that receives remove calls via GRPC protocol.
//
//	Configuration parameters:
//
//	- dependencies:
//		- endpoint:              override for GRPC Endpoint dependency
//		- controller:            override for Controller dependency
//	- connection(s):
//		- discovery_key:         (optional) a key to retrieve the connection from connect.idiscovery.html IDiscovery
//		- protocol:              connection protocol: http or https
//		- host:                  host name or IP address
//		- port:                  port number
//		- uri:                   resource URI or connection string with all parameters in it
//	- credential - the HTTPS credentials:
//		- ssl_key_file:         the SSL private key in PEM
//		- ssl_crt_file:         the SSL certificate in PEM
//		- ssl_ca_file:          the certificate authorities (root cerfiticates) in PEM
//
//	References:
//
//		- *:logger:*:*:1.0               (optional) ILogger components to pass log messages
//		- *:counters:*:*:1.0             (optional) ICounters components to pass collected measurements
//		- *:discovery:*:*:1.0            (optional) IDiscovery services to resolve connection
//		- *:endpoint:grpc:*:1.0           (optional) GrpcEndpoint reference
//
// See GrpcClient
//
// Example:
//
//	type MyGrpcService struct{
//	   *GrpcService
//	   controller IMyController;
//	}
//	...
//
//	func NewMyGrpcService() *MyGrpcService {
//		c := NewMyGrpcService{}
//		c.GrpcService = grpcservices.NewGrpcService("Mydata.Mydatas")
//		c.GrpcService.IRegisterable = &c
//		c.numberOfCalls = 0
//		c.DependencyResolver.Put(context.Context(), "controller", cref.NewDescriptor("mygroup", "controller", "*", "*", "*"))
//		return &c
//	}
//
//	func (c*MyGrpcService) SetReferences(ctx context.Context, references: IReferences) {
//		c.GrpcService.SetReferences(references);
//		resolv, err := c.DependencyResolver.GetOneRequired("controller")
//		if err == nil && resolv != nil {
//		    c.controller = resolv.(grpctest.IMyController)
//		    return
//		}
//		panic("Can't resolve 'controller' reference")
//	}
//
//	func (c*MyGrpcService) Register() {
//		protos.RegisterMyDataServer(c.Endpoint.GetServer(), c)
//		...
//	}
//
//	service := NewMyGrpcService();
//	service.Configure(ctx, cconf.NewConfigParamsFromTuples(
//	    "connection.protocol", "http",
//	    "connection.host", "localhost",
//	    "connection.port", 8080,
//	));
//	service.SetReferences(ctx, cref.NewReferencesFromTuples(
//	   cref.NewDescriptor("mygroup","controller","default","default","1.0"), controller
//	));
//
//	err := service.Open(ctx, "123")
//	if  err == nil {
//	   fmt.Println("The GRPC service is running on port 8080");
//	}
//
type GrpcService struct {
	Overrides IGrpcServiceOverrides

	defaultConfig *cconf.ConfigParams
	serviceName   string
	config        *cconf.ConfigParams
	references    cref.IReferences
	localEndpoint bool
	opened        bool
	//  The GRPC endpoint that exposes c service.
	Endpoint *GrpcEndpoint
	//  The dependency resolver.
	DependencyResolver *cref.DependencyResolver
	//  The logger.
	Logger *clog.CompositeLogger
	//  The performance counters.
	Counters *ccount.CompositeCounters
	// The tracer.
	Tracer *ctrace.CompositeTracer
}

// InheritGrpcService methods are creates new instance NewGrpcService
// Parameters:
//    - overrides a reference to child class that overrides virtual methods
//    - serviceName string
// service name from XYZ.pb.go, set "" for use default gRPC commandable protobuf
// Return *GrpcService
func InheritGrpcService(overrides IGrpcServiceOverrides, serviceName string) *GrpcService {
	c := &GrpcService{
		Overrides: overrides,
	}
	c.serviceName = serviceName
	c.defaultConfig = cconf.NewConfigParamsFromTuples(
		"dependencies.endpoint", "*:endpoint:grpc:*:1.0",
	)
	c.DependencyResolver = cref.NewDependencyResolverWithParams(context.Background(), c.defaultConfig, cref.NewEmptyReferences())
	c.Logger = clog.NewCompositeLogger()
	c.Counters = ccount.NewCompositeCounters()
	c.Tracer = ctrace.NewCompositeTracer()
	return c
}

// Configure method are configures component by passing configuration parameters.
//	Parameters:
//		- ctx context.Context	operation context
//		- config   configuration parameters to be set.
func (c *GrpcService) Configure(ctx context.Context, config *cconf.ConfigParams) {
	config = config.SetDefaults(c.defaultConfig)
	c.config = config
	c.DependencyResolver.Configure(ctx, config)
}

//SetReferences method are sets references to dependent components.
//	Parameters:
//		- ctx context.Context	operation context
//		- references 	references to locate the component dependencies.
func (c *GrpcService) SetReferences(ctx context.Context, references cref.IReferences) {
	c.references = references
	c.Logger.SetReferences(ctx, references)
	c.Counters.SetReferences(ctx, references)
	c.Tracer.SetReferences(ctx, references)
	c.DependencyResolver.SetReferences(ctx, references)
	// Get endpoint
	res := c.DependencyResolver.GetOneOptional("endpoint")
	c.Endpoint, _ = res.(*GrpcEndpoint)
	// Or create a local one
	if c.Endpoint == nil {
		c.Endpoint = c.createEndpoint()
		c.localEndpoint = true
	} else {
		c.localEndpoint = false
	}
	// Add registration callback to the endpoint
	c.Endpoint.Register(c)
}

// UnsetReferences method are unsets (clears) previously set references to dependent components.
func (c *GrpcService) UnsetReferences() {
	// Remove registration callback from endpoint
	if c.Endpoint != nil {
		c.Endpoint.Unregister(c)
		c.Endpoint = nil
	}
}

func (c *GrpcService) createEndpoint() *GrpcEndpoint {
	endpoint := NewGrpcEndpoint()
	if c.config != nil {
		endpoint.Configure(context.Background(), c.config)
	}
	if c.references != nil {
		endpoint.SetReferences(context.Background(), c.references)
	}
	return endpoint
}

// Instrument method are adds instrumentation to log calls and measure call time.
// It returns a Timing object that is used to end the time measurement.
//	Parameters:
//		- ctx context.Context	operation context
//		- correlationId     (optional) transaction id to trace execution through call chain.
//		- name              a method name.
// Return Timing object to end the time measurement.
func (c *GrpcService) Instrument(ctx context.Context, correlationId string, name string) *rpcserv.InstrumentTiming {
	c.Logger.Trace(ctx, correlationId, "Executing %s method", name)
	c.Counters.IncrementOne(ctx, name+".exec_count")

	counterTiming := c.Counters.BeginTiming(ctx, name+".exec_time")
	traceTiming := c.Tracer.BeginTrace(ctx, correlationId, name, "")
	return rpcserv.NewInstrumentTiming(correlationId, name, "exec",
		c.Logger, c.Counters, counterTiming, traceTiming)
}

// InstrumentError method are adds instrumentation to error handling.
//	Parameters:
//		- ctx context.Context	operation context
//		- correlationId     (optional) transaction id to trace execution through call chain.
//		- name              a method name.
//		- errIn               an occured error
//		- resIn            (optional) an execution result
// Returns: result any, err error
// input result and error
// func (c *GrpcService) InstrumentError(correlationId string, name string, errIn error,
// 	resIn any) (result any, err error) {
// 	if errIn != nil {
// 		c.Logger.Error(correlationId, errIn, "Failed to execute %s method", name)
// 		c.Counters.IncrementOne(name + ".exec_errors")
// 	}
// 	return resIn, errIn
// }

// IsOpen method are checks if the component is opened.
// Return true if the component has been opened and false otherwise.
func (c *GrpcService) IsOpen() bool {
	return c.opened
}

// Open method are opens the component.
//	Parameters:
//		- ctx context.Context	operation context
//		- correlationId 	(optional) transaction id to trace execution through call chain.
// Returns: error or nil no errors occured.
func (c *GrpcService) Open(ctx context.Context, correlationId string) (err error) {
	if c.opened {
		return nil
	}

	if c.Endpoint == nil {
		c.Endpoint = c.createEndpoint()
		c.Endpoint.Register(c)
		c.localEndpoint = true
	}

	if c.localEndpoint {
		opnErr := c.Endpoint.Open(ctx, correlationId)
		if opnErr != nil {
			c.opened = false
			return opnErr
		}
	}
	c.opened = true
	return nil
}

// Close method are closes component and frees used resources.
//	Parameters:
//		- ctx context.Context	operation context
//		- correlationId 	(optional) transaction id to trace execution through call chain.
// Returns: error or nil no errors occured.
func (c *GrpcService) Close(ctx context.Context, correlationId string) (err error) {
	if !c.opened {
		return nil
	}

	if c.Endpoint == nil {
		return cerr.NewInvalidStateError(correlationId, "NO_Endpoint", "HTTP endpoint is missing")
	}

	if c.localEndpoint {
		clsErr := c.Endpoint.Close(ctx, correlationId)
		if clsErr != nil {
			c.opened = false
			return clsErr
		}
	}
	c.opened = false
	return nil
}

// RegisterCommandableMethod method are registers a commandable method in c objects GRPC server (service) by the given name.,
//	Parameters:
//		- ctx context.Context	operation context
//		- method        the GRPC method name.
//		- schema        the schema to use for parameter validation.
//		- action        the action to perform at the given route.
func (c *GrpcService) RegisterCommandableMethod(method string, schema *cvalid.Schema,
	action func(ctx context.Context, correlationId string, data *crun.Parameters) (result any, err error)) {
	c.Endpoint.RegisterCommandableMethod(method, schema, action)
}

// Registers a middleware for methods in GRPC endpoint.
//	Parameters:
//		- action        an action function that is called when middleware is invoked.
func (c *GrpcService) RegisterUnaryInterceptor(action func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error)) {
	if c.Endpoint == nil {
		return
	}

	c.Endpoint.AddInterceptors(grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if strings.HasPrefix(info.FullMethod, "/"+c.serviceName+"/") {
			return action(ctx, req, info, handler)
		}
		return handler(ctx, req)
	}))
}

// Register method are registers all service routes in HTTP endpoint.
func (c *GrpcService) Register() {
	// Override in child classes
	c.Overrides.Register()
}

func (c *GrpcService) ValidateRequest(request any, schema *cvalid.Schema) error {

	buf, err := json.Marshal(request)
	if err != nil {
		return err
	}

	validateObj := make(map[string]any)
	err = json.Unmarshal(buf, &validateObj)
	if err != nil {
		return err
	}

	validateErr := schema.ValidateAndReturnError("", validateObj, false)
	if validateErr != nil {
		return validateErr
	}
	return nil
}
