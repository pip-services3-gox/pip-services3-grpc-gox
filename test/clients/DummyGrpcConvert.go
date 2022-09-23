package test_clients

import (
	"encoding/json"

	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	tdata "github.com/pip-services3-gox/pip-services3-grpc-gox/test/data"
	testproto "github.com/pip-services3-gox/pip-services3-grpc-gox/test/protos"
)

func fromError(err error) *testproto.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &testproto.ErrorDescription{
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *testproto.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}
	return r
}

func toMap(val map[string]string) map[string]any {
	r := make(map[string]any)

	for k, v := range val {
		r[k] = v
	}
	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromDummy(in *tdata.Dummy) *testproto.Dummy {
	if in == nil {
		return nil
	}

	obj := &testproto.Dummy{
		Id:      in.Id,
		Key:     in.Key,
		Content: in.Content,
	}
	return obj
}

func toDummy(obj *testproto.Dummy) *tdata.Dummy {
	if obj == nil {
		return nil
	}
	dummy := &tdata.Dummy{
		Id:      obj.Id,
		Key:     obj.Key,
		Content: obj.Content,
	}
	return dummy
}

func toDummiesPage(obj *testproto.DummiesPage) *tdata.DummyDataPage {
	if obj == nil {
		return nil
	}

	dummies := make([]tdata.Dummy, len(obj.Data))
	for i, v := range obj.Data {
		dummies[i] = *toDummy(v)
	}
	page := tdata.NewDummyDataPage(&obj.Total, dummies)

	return page
}
