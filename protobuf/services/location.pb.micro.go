// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: protobuf/services/location.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
	ptypes "protobuf/ptypes"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for LocationService service

func NewLocationServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LocationService service

type LocationService interface {
	GetLocation(ctx context.Context, in *wrappers.BytesValue, opts ...client.CallOption) (*ptypes.Location, error)
}

type locationService struct {
	c    client.Client
	name string
}

func NewLocationService(name string, c client.Client) LocationService {
	return &locationService{
		c:    c,
		name: name,
	}
}

func (c *locationService) GetLocation(ctx context.Context, in *wrappers.BytesValue, opts ...client.CallOption) (*ptypes.Location, error) {
	req := c.c.NewRequest(c.name, "LocationService.GetLocation", in)
	out := new(ptypes.Location)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LocationService service

type LocationServiceHandler interface {
	GetLocation(context.Context, *wrappers.BytesValue, *ptypes.Location) error
}

func RegisterLocationServiceHandler(s server.Server, hdlr LocationServiceHandler, opts ...server.HandlerOption) error {
	type locationService interface {
		GetLocation(ctx context.Context, in *wrappers.BytesValue, out *ptypes.Location) error
	}
	type LocationService struct {
		locationService
	}
	h := &locationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&LocationService{h}, opts...))
}

type locationServiceHandler struct {
	LocationServiceHandler
}

func (h *locationServiceHandler) GetLocation(ctx context.Context, in *wrappers.BytesValue, out *ptypes.Location) error {
	return h.LocationServiceHandler.GetLocation(ctx, in, out)
}
