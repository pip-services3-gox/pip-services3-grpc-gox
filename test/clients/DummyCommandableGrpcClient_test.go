package test_clients

import (
	"context"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	tlogic "github.com/pip-services3-gox/pip-services3-grpc-gox/test/logic"
	testservices "github.com/pip-services3-gox/pip-services3-grpc-gox/test/services"
)

func TestDummyRestClient(t *testing.T) {
	ctx := context.Background()

	grpcConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", "localhost",
		"connection.port", "3002",
	)

	var service *testservices.DummyCommandableGrpcService
	var client *DummyCommandableGrpcClient
	var fixture *DummyClientFixture

	ctrl := tlogic.NewDummyController()

	service = testservices.NewDummyCommandableGrpcService()
	service.Configure(ctx, grpcConfig)

	references := cref.NewReferencesFromTuples(ctx,
		cref.NewDescriptor("pip-services-dummies", "controller", "default", "default", "1.0"), ctrl,
		cref.NewDescriptor("pip-services-dummies", "service", "grpc", "default", "1.0"), service,
	)
	service.SetReferences(ctx, references)

	service.Open(ctx, "")
	defer service.Close(ctx, "")

	client = NewDummyCommandableGrpcClient()
	fixture = NewDummyClientFixture(client)

	client.Configure(ctx, grpcConfig)
	client.SetReferences(ctx, cref.NewEmptyReferences())
	client.Open(ctx, "")
	defer client.Close(ctx, "")

	t.Run("CRUD Operations", fixture.TestCrudOperations)

}
