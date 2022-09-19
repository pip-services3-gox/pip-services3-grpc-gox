package test_clients

import (
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	tlogic "github.com/pip-services3-gox/pip-services3-grpc-gox/test/logic"
	testservices "github.com/pip-services3-gox/pip-services3-grpc-gox/test/services"
)

func TestDummyRestClient(t *testing.T) {
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
	service.Configure(grpcConfig)

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services-dummies", "controller", "default", "default", "1.0"), ctrl,
		cref.NewDescriptor("pip-services-dummies", "service", "grpc", "default", "1.0"), service,
	)
	service.SetReferences(references)

	service.Open("")

	defer service.Close("")

	client = NewDummyCommandableGrpcClient()
	fixture = NewDummyClientFixture(client)

	client.Configure(grpcConfig)
	client.SetReferences(cref.NewEmptyReferences())
	client.Open("")
	defer client.Close("")

	t.Run("CRUD Operations", fixture.TestCrudOperations)

}
