package test_logic

import (
	"context"
	"encoding/json"

	ccomand "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	crun "github.com/pip-services3-gox/pip-services3-commons-gox/run"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	tdata "github.com/pip-services3-gox/pip-services3-grpc-gox/test/data"
)

type DummyCommandSet struct {
	ccomand.CommandSet
	controller IDummyController
}

func NewDummyCommandSet(controller IDummyController) *DummyCommandSet {
	dcs := DummyCommandSet{}
	dcs.CommandSet = *ccomand.NewCommandSet()

	dcs.controller = controller

	dcs.AddCommand(dcs.makeGetPageByFilterCommand())
	dcs.AddCommand(dcs.makeGetOneByIdCommand())
	dcs.AddCommand(dcs.makeCreateCommand())
	dcs.AddCommand(dcs.makeUpdateCommand())
	dcs.AddCommand(dcs.makeDeleteByIdCommand())
	return &dcs
}

func (c *DummyCommandSet) makeGetPageByFilterCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"get_dummies",
		cvalid.NewObjectSchema().WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			var filter *cdata.FilterParams
			var paging *cdata.PagingParams

			if _val, ok := args.Get("filter"); ok {
				filter = cdata.NewFilterParamsFromValue(_val)
			}
			if _val, ok := args.Get("paging"); ok {
				paging = cdata.NewPagingParamsFromValue(_val)
			}

			return c.controller.GetPageByFilter(correlationId, filter, paging)
		},
	)
}

func (c *DummyCommandSet) makeGetOneByIdCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"get_dummy_by_id",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy_id", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			id := args.GetAsString("dummy_id")
			return c.controller.GetOneById(correlationId, id)
		},
	)
}

func (c *DummyCommandSet) makeCreateCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"create_dummy",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy", tdata.NewDummySchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			var entity tdata.Dummy

			if _val, ok := args.Get("dummy"); ok {
				val, _ := json.Marshal(_val)
				json.Unmarshal(val, &entity)
			}

			return c.controller.Create(correlationId, entity)
		},
	)
}

func (c *DummyCommandSet) makeUpdateCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"update_dummy",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy", tdata.NewDummySchema()),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			var entity tdata.Dummy

			if _val, ok := args.Get("dummy"); ok {
				val, _ := json.Marshal(_val)
				json.Unmarshal(val, &entity)
			}
			return c.controller.Update(correlationId, entity)
		},
	)
}

func (c *DummyCommandSet) makeDeleteByIdCommand() ccomand.ICommand {
	return ccomand.NewCommand(
		"delete_dummy_by_id",
		cvalid.NewObjectSchema().WithRequiredProperty("dummy_id", cconv.String),
		func(ctx context.Context, correlationId string, args *crun.Parameters) (result any, err error) {
			id := args.GetAsString("dummy_id")
			return c.controller.DeleteById(correlationId, id)
		},
	)
}
