// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/elastic/elastic.proto

/*
Package go_micro_srv_elastic is a generated protocol buffer package.

It is generated from these files:
	proto/elastic/elastic.proto

It has these top-level messages:
	DocRef
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
	QueryRequest
	QueryResponse
	CountRequest
	CountResponse
	CreateIndexWithSettingsRequest
	CreateIndexWithSettingsResponse
	PutMappingFromJSONRequest
	PutMappingFromJSONResponse
*/
package go_micro_srv_elastic

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Elastic service

type ElasticService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
	CreateIndexWithSettings(ctx context.Context, in *CreateIndexWithSettingsRequest, opts ...client.CallOption) (*CreateIndexWithSettingsResponse, error)
	PutMappingFromJSON(ctx context.Context, in *PutMappingFromJSONRequest, opts ...client.CallOption) (*PutMappingFromJSONRequest, error)
}

type elasticService struct {
	c    client.Client
	name string
}

func NewElasticService(name string, c client.Client) ElasticService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.elastic"
	}
	return &elasticService{
		c:    c,
		name: name,
	}
}

func (c *elasticService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Elastic.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Elastic.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Elastic.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Elastic.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.name, "Elastic.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticService) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.name, "Elastic.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticService) CreateIndexWithSettings(ctx context.Context, in *CreateIndexWithSettingsRequest, opts ...client.CallOption) (*CreateIndexWithSettingsResponse, error) {
	req := c.c.NewRequest(c.name, "Elastic.CreateIndexWithSettings", in)
	out := new(CreateIndexWithSettingsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticService) PutMappingFromJSON(ctx context.Context, in *PutMappingFromJSONRequest, opts ...client.CallOption) (*PutMappingFromJSONRequest, error) {
	req := c.c.NewRequest(c.name, "Elastic.PutMappingFromJSON", in)
	out := new(PutMappingFromJSONRequest)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Elastic service

type ElasticHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	Query(context.Context, *QueryRequest, *QueryResponse) error
	CreateIndexWithSettings(context.Context, *CreateIndexWithSettingsRequest, *CreateIndexWithSettingsResponse) error
	PutMappingFromJSON(context.Context, *PutMappingFromJSONRequest, *PutMappingFromJSONRequest) error
}

func RegisterElasticHandler(s server.Server, hdlr ElasticHandler, opts ...server.HandlerOption) error {
	type elastic interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error
		Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error
		CreateIndexWithSettings(ctx context.Context, in *CreateIndexWithSettingsRequest, out *CreateIndexWithSettingsResponse) error
		PutMappingFromJSON(ctx context.Context, in *PutMappingFromJSONRequest, out *PutMappingFromJSONRequest) error
	}
	type Elastic struct {
		elastic
	}
	h := &elasticHandler{hdlr}
	return s.Handle(s.NewHandler(&Elastic{h}, opts...))
}

type elasticHandler struct {
	ElasticHandler
}

func (h *elasticHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ElasticHandler.Create(ctx, in, out)
}

func (h *elasticHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ElasticHandler.Read(ctx, in, out)
}

func (h *elasticHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ElasticHandler.Update(ctx, in, out)
}

func (h *elasticHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ElasticHandler.Delete(ctx, in, out)
}

func (h *elasticHandler) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.ElasticHandler.Search(ctx, in, out)
}

func (h *elasticHandler) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.ElasticHandler.Query(ctx, in, out)
}

func (h *elasticHandler) CreateIndexWithSettings(ctx context.Context, in *CreateIndexWithSettingsRequest, out *CreateIndexWithSettingsResponse) error {
	return h.ElasticHandler.CreateIndexWithSettings(ctx, in, out)
}

func (h *elasticHandler) PutMappingFromJSON(ctx context.Context, in *PutMappingFromJSONRequest, out *PutMappingFromJSONRequest) error {
	return h.ElasticHandler.PutMappingFromJSON(ctx, in, out)
}