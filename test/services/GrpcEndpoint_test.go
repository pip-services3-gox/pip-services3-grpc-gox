package test_services

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	grpcservices "github.com/pip-services3-gox/pip-services3-grpc-gox/services"
	"github.com/stretchr/testify/assert"
)

func TestGrpcEndpoint(t *testing.T) {
	ctx := context.Background()

	grpcConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", "localhost",
		"connection.port", 3005,
	)

	endpoint := grpcservices.NewGrpcEndpoint()
	endpoint.Configure(ctx, grpcConfig)

	endpoint.Open(ctx, "")
	assert.True(t, endpoint.IsOpen())
	endpoint.Close(ctx, "")
}
