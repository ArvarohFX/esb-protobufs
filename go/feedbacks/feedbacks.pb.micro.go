// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/feedbacks.proto

package feedbacks

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for Mobile service

func NewMobileEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Mobile service

type MobileService interface {
	App(ctx context.Context, in *ParamsApp, opts ...client.CallOption) (*ResponseOk, error)
	Store(ctx context.Context, in *ParamsStore, opts ...client.CallOption) (*ResponseOk, error)
	Order(ctx context.Context, in *ParamsOrder, opts ...client.CallOption) (*ResponseOk, error)
	Categories(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseCategories, error)
	ReasonsByOrder(ctx context.Context, in *ParamsReasonsByOrder, opts ...client.CallOption) (*ResponseReasons, error)
	ReasonsByStore(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseReasons, error)
	CanBeSaved(ctx context.Context, in *CanBeSavedParams, opts ...client.CallOption) (*ResponseOk, error)
}

type mobileService struct {
	c    client.Client
	name string
}

func NewMobileService(name string, c client.Client) MobileService {
	return &mobileService{
		c:    c,
		name: name,
	}
}

func (c *mobileService) App(ctx context.Context, in *ParamsApp, opts ...client.CallOption) (*ResponseOk, error) {
	req := c.c.NewRequest(c.name, "Mobile.App", in)
	out := new(ResponseOk)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) Store(ctx context.Context, in *ParamsStore, opts ...client.CallOption) (*ResponseOk, error) {
	req := c.c.NewRequest(c.name, "Mobile.Store", in)
	out := new(ResponseOk)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) Order(ctx context.Context, in *ParamsOrder, opts ...client.CallOption) (*ResponseOk, error) {
	req := c.c.NewRequest(c.name, "Mobile.Order", in)
	out := new(ResponseOk)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) Categories(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseCategories, error) {
	req := c.c.NewRequest(c.name, "Mobile.Categories", in)
	out := new(ResponseCategories)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) ReasonsByOrder(ctx context.Context, in *ParamsReasonsByOrder, opts ...client.CallOption) (*ResponseReasons, error) {
	req := c.c.NewRequest(c.name, "Mobile.ReasonsByOrder", in)
	out := new(ResponseReasons)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) ReasonsByStore(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ResponseReasons, error) {
	req := c.c.NewRequest(c.name, "Mobile.ReasonsByStore", in)
	out := new(ResponseReasons)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileService) CanBeSaved(ctx context.Context, in *CanBeSavedParams, opts ...client.CallOption) (*ResponseOk, error) {
	req := c.c.NewRequest(c.name, "Mobile.CanBeSaved", in)
	out := new(ResponseOk)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mobile service

type MobileHandler interface {
	App(context.Context, *ParamsApp, *ResponseOk) error
	Store(context.Context, *ParamsStore, *ResponseOk) error
	Order(context.Context, *ParamsOrder, *ResponseOk) error
	Categories(context.Context, *empty.Empty, *ResponseCategories) error
	ReasonsByOrder(context.Context, *ParamsReasonsByOrder, *ResponseReasons) error
	ReasonsByStore(context.Context, *empty.Empty, *ResponseReasons) error
	CanBeSaved(context.Context, *CanBeSavedParams, *ResponseOk) error
}

func RegisterMobileHandler(s server.Server, hdlr MobileHandler, opts ...server.HandlerOption) error {
	type mobile interface {
		App(ctx context.Context, in *ParamsApp, out *ResponseOk) error
		Store(ctx context.Context, in *ParamsStore, out *ResponseOk) error
		Order(ctx context.Context, in *ParamsOrder, out *ResponseOk) error
		Categories(ctx context.Context, in *empty.Empty, out *ResponseCategories) error
		ReasonsByOrder(ctx context.Context, in *ParamsReasonsByOrder, out *ResponseReasons) error
		ReasonsByStore(ctx context.Context, in *empty.Empty, out *ResponseReasons) error
		CanBeSaved(ctx context.Context, in *CanBeSavedParams, out *ResponseOk) error
	}
	type Mobile struct {
		mobile
	}
	h := &mobileHandler{hdlr}
	return s.Handle(s.NewHandler(&Mobile{h}, opts...))
}

type mobileHandler struct {
	MobileHandler
}

func (h *mobileHandler) App(ctx context.Context, in *ParamsApp, out *ResponseOk) error {
	return h.MobileHandler.App(ctx, in, out)
}

func (h *mobileHandler) Store(ctx context.Context, in *ParamsStore, out *ResponseOk) error {
	return h.MobileHandler.Store(ctx, in, out)
}

func (h *mobileHandler) Order(ctx context.Context, in *ParamsOrder, out *ResponseOk) error {
	return h.MobileHandler.Order(ctx, in, out)
}

func (h *mobileHandler) Categories(ctx context.Context, in *empty.Empty, out *ResponseCategories) error {
	return h.MobileHandler.Categories(ctx, in, out)
}

func (h *mobileHandler) ReasonsByOrder(ctx context.Context, in *ParamsReasonsByOrder, out *ResponseReasons) error {
	return h.MobileHandler.ReasonsByOrder(ctx, in, out)
}

func (h *mobileHandler) ReasonsByStore(ctx context.Context, in *empty.Empty, out *ResponseReasons) error {
	return h.MobileHandler.ReasonsByStore(ctx, in, out)
}

func (h *mobileHandler) CanBeSaved(ctx context.Context, in *CanBeSavedParams, out *ResponseOk) error {
	return h.MobileHandler.CanBeSaved(ctx, in, out)
}

// Api Endpoints for Store service

func NewStoreEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Store service

type StoreService interface {
	New(ctx context.Context, in *NewParams, opts ...client.CallOption) (*NewResponse, error)
	Patch(ctx context.Context, in *PatchParams, opts ...client.CallOption) (*ResponseOk, error)
}

type storeService struct {
	c    client.Client
	name string
}

func NewStoreService(name string, c client.Client) StoreService {
	return &storeService{
		c:    c,
		name: name,
	}
}

func (c *storeService) New(ctx context.Context, in *NewParams, opts ...client.CallOption) (*NewResponse, error) {
	req := c.c.NewRequest(c.name, "Store.New", in)
	out := new(NewResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Patch(ctx context.Context, in *PatchParams, opts ...client.CallOption) (*ResponseOk, error) {
	req := c.c.NewRequest(c.name, "Store.Patch", in)
	out := new(ResponseOk)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreHandler interface {
	New(context.Context, *NewParams, *NewResponse) error
	Patch(context.Context, *PatchParams, *ResponseOk) error
}

func RegisterStoreHandler(s server.Server, hdlr StoreHandler, opts ...server.HandlerOption) error {
	type store interface {
		New(ctx context.Context, in *NewParams, out *NewResponse) error
		Patch(ctx context.Context, in *PatchParams, out *ResponseOk) error
	}
	type Store struct {
		store
	}
	h := &storeHandler{hdlr}
	return s.Handle(s.NewHandler(&Store{h}, opts...))
}

type storeHandler struct {
	StoreHandler
}

func (h *storeHandler) New(ctx context.Context, in *NewParams, out *NewResponse) error {
	return h.StoreHandler.New(ctx, in, out)
}

func (h *storeHandler) Patch(ctx context.Context, in *PatchParams, out *ResponseOk) error {
	return h.StoreHandler.Patch(ctx, in, out)
}
