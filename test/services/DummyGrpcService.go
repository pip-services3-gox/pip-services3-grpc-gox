package test_services

import (
	"context"
	"encoding/json"

	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	grpcservices "github.com/pip-services3-gox/pip-services3-grpc-gox/services"
	tdata "github.com/pip-services3-gox/pip-services3-grpc-gox/test/data"
	tlogic "github.com/pip-services3-gox/pip-services3-grpc-gox/test/logic"
	"google.golang.org/grpc"

	grpcproto "github.com/pip-services3-gox/pip-services3-grpc-gox/test/protos"
)

type DummyGrpcService struct {
	*grpcservices.GrpcService
	controller    tlogic.IDummyController
	numberOfCalls int64

	grpcproto.DummiesServer
}

func NewDummyGrpcService() *DummyGrpcService {
	c := &DummyGrpcService{}
	c.GrpcService = grpcservices.InheritGrpcService(c, "dummies.Dummies")
	c.numberOfCalls = 0
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("pip-services-dummies", "controller", "default", "*", "*"))
	return c
}

func (c *DummyGrpcService) SetReferences(ctx context.Context, references cref.IReferences) {
	c.GrpcService.SetReferences(ctx, references)
	resolv, err := c.DependencyResolver.GetOneRequired("controller")
	if err == nil && resolv != nil {
		c.controller = resolv.(tlogic.IDummyController)
		return
	}
	panic("Can't resolve 'controller' reference")
}

func (c *DummyGrpcService) GetNumberOfCalls() int64 {
	return c.numberOfCalls
}

func (c *DummyGrpcService) incrementNumberOfCalls(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {

	m, err := handler(ctx, req)
	if err != nil {
		c.Logger.Error(ctx, "", err, "RPC failed with error %v", err.Error())
	}
	c.numberOfCalls++
	return m, err
}

func (c *DummyGrpcService) Open(ctx context.Context, correlationId string) error {

	// Add interceptors
	c.RegisterUnaryInterceptor(c.incrementNumberOfCalls)
	return c.GrpcService.Open(ctx, correlationId)
}

func (c *DummyGrpcService) GetDummies(ctx context.Context, req *grpcproto.DummiesPageRequest) (*grpcproto.DummiesPage, error) {

	validateErr := c.ValidateRequest(req,
		&cvalid.NewObjectSchema().
			WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()).
			WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).Schema)

	if validateErr != nil {
		return nil, validateErr
	}

	filter := cdata.NewFilterParamsFromValue(req.GetFilter())
	paging := cdata.NewEmptyPagingParams()
	if req.Paging != nil {
		paging = cdata.NewPagingParams(req.Paging.GetSkip(), req.Paging.GetTake(), req.Paging.GetTotal())
	}
	data, err := c.controller.GetPageByFilter(
		req.CorrelationId,
		filter,
		paging,
	)
	if err != nil || data == nil {
		return nil, err
	}

	result := grpcproto.DummiesPage{}
	result.Total = *data.Total
	for _, v := range data.Data {
		buf := grpcproto.Dummy{}
		bytes, _ := json.Marshal(v)
		json.Unmarshal(bytes, &buf)
		result.Data = append(result.Data, &buf)
	}
	return &result, err
}

func (c *DummyGrpcService) GetDummyById(ctx context.Context, req *grpcproto.DummyIdRequest) (*grpcproto.Dummy, error) {

	// validation
	validateErr := c.ValidateRequest(req,
		&cvalid.NewObjectSchema().
			WithRequiredProperty("dummy_id", cconv.String).Schema)

	if validateErr != nil {
		return nil, validateErr
	}
	// ==================================

	data, err := c.controller.GetOneById(
		req.CorrelationId,
		req.DummyId,
	)
	if err != nil {
		return nil, err
	}
	result := grpcproto.Dummy{}
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &result)
	return &result, nil
}

func (c *DummyGrpcService) CreateDummy(ctx context.Context, req *grpcproto.DummyObjectRequest) (*grpcproto.Dummy, error) {

	// validation
	validateErr := c.ValidateRequest(req,
		&cvalid.NewObjectSchema().
			WithRequiredProperty("dummy", tdata.NewDummySchema()).Schema)

	if validateErr != nil {
		return nil, validateErr
	}

	dummy := tdata.Dummy{}
	bytes, _ := json.Marshal(req.Dummy)
	json.Unmarshal(bytes, &dummy)

	data, err := c.controller.Create(
		req.CorrelationId,
		dummy,
	)

	if err != nil || data == nil {
		return nil, err
	}
	result := grpcproto.Dummy{}
	bytes, _ = json.Marshal(data)
	json.Unmarshal(bytes, &result)
	return &result, nil
}

func (c *DummyGrpcService) UpdateDummy(ctx context.Context, req *grpcproto.DummyObjectRequest) (*grpcproto.Dummy, error) {

	validateErr := c.ValidateRequest(req,
		&cvalid.NewObjectSchema().
			WithRequiredProperty("dummy", tdata.NewDummySchema()).Schema)

	if validateErr != nil {
		return nil, validateErr
	}

	dummy := tdata.Dummy{}
	bytes, _ := json.Marshal(req.Dummy)
	json.Unmarshal(bytes, &dummy)

	data, err := c.controller.Update(
		req.CorrelationId,
		dummy,
	)

	if err != nil || data == nil {
		return nil, err
	}
	result := grpcproto.Dummy{}
	bytes, _ = json.Marshal(data)
	json.Unmarshal(bytes, &result)
	return &result, nil
}

func (c *DummyGrpcService) DeleteDummyById(ctx context.Context, req *grpcproto.DummyIdRequest) (*grpcproto.Dummy, error) {

	validateErr := c.ValidateRequest(req,
		&cvalid.NewObjectSchema().
			WithRequiredProperty("dummy_id", cconv.String).Schema)

	if validateErr != nil {
		return nil, validateErr
	}

	data, err := c.controller.DeleteById(
		req.CorrelationId,
		req.DummyId,
	)
	if err != nil || data == nil {
		return nil, err
	}
	result := grpcproto.Dummy{}
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &result)
	return &result, nil
}

func (c *DummyGrpcService) Register() {
	grpcproto.RegisterDummiesServer(c.Endpoint.GetServer(), c)
}
