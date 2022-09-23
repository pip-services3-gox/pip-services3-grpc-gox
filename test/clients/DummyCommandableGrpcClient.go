package test_clients

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	grpcclients "github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
	tdata "github.com/pip-services3-gox/pip-services3-grpc-gox/test/data"
)

type DummyCommandableGrpcClient struct {
	*grpcclients.CommandableGrpcClient
}

func NewDummyCommandableGrpcClient() *DummyCommandableGrpcClient {
	dcgc := DummyCommandableGrpcClient{}
	dcgc.CommandableGrpcClient = grpcclients.NewCommandableGrpcClient("dummy")
	return &dcgc
}

func (c *DummyCommandableGrpcClient) GetDummies(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result *tdata.DummyDataPage, err error) {

	params := cdata.NewEmptyStringValueMap()
	c.AddFilterParams(params, filter)
	c.AddPagingParams(params, paging)

	response, calErr := c.CallCommand(ctx, "get_dummies", correlationId, cdata.NewAnyValueMapFromValue(params.Value()))
	if calErr != nil {
		return nil, calErr
	}

	return grpcclients.HandleHttpResponse[*tdata.DummyDataPage](response, correlationId)
}

func (c *DummyCommandableGrpcClient) GetDummyById(ctx context.Context, correlationId string, dummyId string) (result *tdata.Dummy, err error) {

	params := cdata.NewEmptyAnyValueMap()
	params.Put("dummy_id", dummyId)

	response, calErr := c.CallCommand(ctx, "get_dummy_by_id", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}

	return grpcclients.HandleHttpResponse[*tdata.Dummy](response, correlationId)
}

func (c *DummyCommandableGrpcClient) CreateDummy(ctx context.Context, correlationId string, dummy tdata.Dummy) (result *tdata.Dummy, err error) {

	params := cdata.NewEmptyAnyValueMap()
	params.Put("dummy", dummy)

	response, calErr := c.CallCommand(ctx, "create_dummy", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}

	return grpcclients.HandleHttpResponse[*tdata.Dummy](response, correlationId)
}

func (c *DummyCommandableGrpcClient) UpdateDummy(ctx context.Context, correlationId string, dummy tdata.Dummy) (result *tdata.Dummy, err error) {

	params := cdata.NewEmptyAnyValueMap()
	params.Put("dummy", dummy)

	response, calErr := c.CallCommand(ctx, "update_dummy", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}

	return grpcclients.HandleHttpResponse[*tdata.Dummy](response, correlationId)
}

func (c *DummyCommandableGrpcClient) DeleteDummy(ctx context.Context, correlationId string, dummyId string) (result *tdata.Dummy, err error) {

	params := cdata.NewEmptyAnyValueMap()
	params.Put("dummy_id", dummyId)

	response, calErr := c.CallCommand(ctx, "delete_dummy_by_id", correlationId, params)
	if calErr != nil {
		return nil, calErr
	}

	return grpcclients.HandleHttpResponse[*tdata.Dummy](response, correlationId)
}
