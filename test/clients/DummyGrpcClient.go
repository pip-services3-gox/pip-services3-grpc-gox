package test_clients

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	grpcclients "github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
	tdata "github.com/pip-services3-gox/pip-services3-grpc-gox/test/data"
	testproto "github.com/pip-services3-gox/pip-services3-grpc-gox/test/protos"
)

type DummyGrpcClient struct {
	*grpcclients.GrpcClient
}

func NewDummyGrpcClient() *DummyGrpcClient {
	dgc := DummyGrpcClient{}
	dgc.GrpcClient = grpcclients.NewGrpcClient("dummies.Dummies")
	return &dgc
}

func (c *DummyGrpcClient) GetDummies(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result *tdata.DummyDataPage, err error) {

	req := &testproto.DummiesPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &testproto.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  paging.GetTake(100),
			Total: paging.Total,
		}
	}
	reply := new(testproto.DummiesPage)
	err = c.CallWithContext(ctx, correlationId, "get_dummies", req, reply)
	c.Instrument(ctx, correlationId, "dummy.get_page_by_filter")
	if err != nil {
		return nil, err
	}
	result = toDummiesPage(reply)
	return result, nil
}

func (c *DummyGrpcClient) GetDummyById(ctx context.Context, correlationId string, dummyId string) (result *tdata.Dummy, err error) {

	req := &testproto.DummyIdRequest{
		CorrelationId: correlationId,
		DummyId:       dummyId,
	}

	reply := new(testproto.Dummy)
	err = c.CallWithContext(ctx, correlationId, "get_dummy_by_id", req, reply)
	c.Instrument(ctx, correlationId, "dummy.get_one_by_id")
	if err != nil {
		return nil, err
	}
	result = toDummy(reply)
	if result != nil && result.Id == "" && result.Key == "" {
		result = nil
	}
	return result, nil
}

func (c *DummyGrpcClient) CreateDummy(ctx context.Context, correlationId string, dummy tdata.Dummy) (result *tdata.Dummy, err error) {

	req := &testproto.DummyObjectRequest{
		CorrelationId: correlationId,
		Dummy:         fromDummy(&dummy),
	}

	reply := new(testproto.Dummy)
	err = c.CallWithContext(ctx, correlationId, "create_dummy", req, reply)
	c.Instrument(ctx, correlationId, "dummy.create")
	if err != nil {
		return nil, err
	}
	result = toDummy(reply)
	if result != nil && result.Id == "" && result.Key == "" {
		result = nil
	}
	return result, nil
}

func (c *DummyGrpcClient) UpdateDummy(ctx context.Context, correlationId string, dummy tdata.Dummy) (result *tdata.Dummy, err error) {
	req := &testproto.DummyObjectRequest{
		CorrelationId: correlationId,
		Dummy:         fromDummy(&dummy),
	}
	reply := new(testproto.Dummy)
	err = c.CallWithContext(ctx, correlationId, "update_dummy", req, reply)
	c.Instrument(ctx, correlationId, "dummy.update")
	if err != nil {
		return nil, err
	}
	result = toDummy(reply)
	if result != nil && result.Id == "" && result.Key == "" {
		result = nil
	}
	return result, nil
}

func (c *DummyGrpcClient) DeleteDummy(ctx context.Context, correlationId string, dummyId string) (result *tdata.Dummy, err error) {

	req := &testproto.DummyIdRequest{
		CorrelationId: correlationId,
		DummyId:       dummyId,
	}

	reply := new(testproto.Dummy)
	c.CallWithContext(ctx, correlationId, "delete_dummy_by_id", req, reply)
	c.Instrument(ctx, correlationId, "dummy.delete_by_id")
	if err != nil {
		return nil, err
	}
	result = toDummy(reply)
	if result != nil && result.Id == "" && result.Key == "" {
		result = nil
	}
	return result, nil
}
