package clients

import (
	"context"
	"encoding/json"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	grpcproto "github.com/pip-services3-gox/pip-services3-grpc-gox/protos"
)

// CommandableGrpcClient abstract client that calls commandable GRPC service.
//
// Commandable services are generated automatically for ICommandable objects.
// Each command is exposed as Invoke method that receives all parameters as args.
//
//	Configuration parameters:
//
//		- connection(s):
//			- discovery_key:         (optional) a key to retrieve the connection from IDiscovery
//			- protocol:              connection protocol: http or https
//			- host:                  host name or IP address
//			- port:                  port number
//			- uri:                   resource URI or connection string with all parameters in it
//		- options:
//			- retries:               number of retries (default: 3)
//			- connect_timeout:       connection timeout in milliseconds (default: 10 sec)
//			- timeout:               invocation timeout in milliseconds (default: 10 sec)
//
//	References:
//
//		- *:logger:*:*:1.0         (optional) ILogger components to pass log messages
//		- *:counters:*:*:1.0         (optional) ICounters components to pass collected measurements
//		- *:discovery:*:*:1.0        (optional) IDiscovery services to resolve connection
//
// Example:
//
//	type MyCommandableGrpcClient struct {
//		*CommandableGrpcClient
//	   	...
//	}
//
//	func (c * MyCommandableGrpcClient) GetData(ctx context.Context, correlationId string, id string) (result *MyData, err error) {
//	   	params := cdata.NewEmptyStringValueMap()
//	   	params.Put("id", id)
//		response, calErr := c.CallCommand(MyDataType, "get_mydata_by_id", correlationId, params)
//		if calErr != nil {
//		    return nil, calErr
//		}
//		return grpcclients.HandleHttpResponse[*MyData](response, correlationId)
//	}
//	...
//
//	client := NewMyCommandableGrpcClient();
//	client.Configure(ctx, cconf.NewConfigParamsFromTuples(
//	    "connection.protocol", "http",
//	    "connection.host", "localhost",
//	    "connection.port", 8080,
//	));
//
//	result, err := client.GetData(ctx, "123", "1")
//	...
//
type CommandableGrpcClient struct {
	*GrpcClient
	//The service name
	Name string
}

// NewCommandableGrpcClient method are creates a new instance of the client.
// Parameters:
//   - name     a service name.
func NewCommandableGrpcClient(name string) *CommandableGrpcClient {
	c := CommandableGrpcClient{}
	c.GrpcClient = NewGrpcClient("commandable.Commandable")
	c.Name = name
	return &c
}

// CallCommand method are calls a remote method via GRPC commadable protocol.
// The call is made via Invoke method and all parameters are sent in args object.
// The complete route to remote method is defined as serviceName + "." + name.
//	Parameters:
//		- ctx context.Context	operation context
//		- name              a name of the command to call.,
//		- correlationId     (optional) transaction id to trace execution through call chain.
//		- params            command parameters.
// Retruns: result or error.
func (c *CommandableGrpcClient) CallCommand(ctx context.Context, name string, correlationId string, params *cdata.AnyValueMap) (result *grpcproto.InvokeReply, err error) {
	method := c.Name + "." + name
	timing := c.Instrument(ctx, correlationId, method)

	var jsonArgs string
	if params != nil {
		jsonRes, err := json.Marshal(params.Value())
		jsonArgs = string(jsonRes)
		if err != nil {
			return result, err
		}
	}

	request := &grpcproto.InvokeRequest{
		Method:        method,
		CorrelationId: correlationId,
		ArgsEmpty:     params == nil,
		ArgsJson:      jsonArgs,
	}

	response := &grpcproto.InvokeReply{}
	err = c.CallWithContext(ctx, "invoke", correlationId, request, response)

	timing.EndTiming(ctx, err)

	// Handle unexpected error
	if err != nil {
		return response, err
	}

	return response, nil
}
