// The MIT License (MIT)
// 
// Copyright (c) 2021 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/service_meta.proto

package apiv1

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// MetaAPIYARPCClient is the YARPC client-side interface for the MetaAPI service.
type MetaAPIYARPCClient interface {
	Health(context.Context, *HealthRequest, ...yarpc.CallOption) (*HealthResponse, error)
}

func newMetaAPIYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) MetaAPIYARPCClient {
	return &_MetaAPIYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.cadence.api.v1.MetaAPI",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewMetaAPIYARPCClient builds a new YARPC client for the MetaAPI service.
func NewMetaAPIYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) MetaAPIYARPCClient {
	return newMetaAPIYARPCClient(clientConfig, nil, options...)
}

// MetaAPIYARPCServer is the YARPC server-side interface for the MetaAPI service.
type MetaAPIYARPCServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

type buildMetaAPIYARPCProceduresParams struct {
	Server      MetaAPIYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildMetaAPIYARPCProcedures(params buildMetaAPIYARPCProceduresParams) []transport.Procedure {
	handler := &_MetaAPIYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.cadence.api.v1.MetaAPI",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Health",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Health,
							NewRequest:  newMetaAPIServiceHealthYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildMetaAPIYARPCProcedures prepares an implementation of the MetaAPI service for YARPC registration.
func BuildMetaAPIYARPCProcedures(server MetaAPIYARPCServer) []transport.Procedure {
	return buildMetaAPIYARPCProcedures(buildMetaAPIYARPCProceduresParams{Server: server})
}

// FxMetaAPIYARPCClientParams defines the input
// for NewFxMetaAPIYARPCClient. It provides the
// paramaters to get a MetaAPIYARPCClient in an
// Fx application.
type FxMetaAPIYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxMetaAPIYARPCClientResult defines the output
// of NewFxMetaAPIYARPCClient. It provides a
// MetaAPIYARPCClient to an Fx application.
type FxMetaAPIYARPCClientResult struct {
	fx.Out

	Client MetaAPIYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxMetaAPIYARPCClient provides a MetaAPIYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    apiv1.NewFxMetaAPIYARPCClient("service-name"),
//    ...
//  )
func NewFxMetaAPIYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxMetaAPIYARPCClientParams) FxMetaAPIYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxMetaAPIYARPCClientResult{
			Client: newMetaAPIYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxMetaAPIYARPCProceduresParams defines the input
// for NewFxMetaAPIYARPCProcedures. It provides the
// paramaters to get MetaAPIYARPCServer procedures in an
// Fx application.
type FxMetaAPIYARPCProceduresParams struct {
	fx.In

	Server      MetaAPIYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxMetaAPIYARPCProceduresResult defines the output
// of NewFxMetaAPIYARPCProcedures. It provides
// MetaAPIYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxMetaAPIYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxMetaAPIYARPCProcedures provides MetaAPIYARPCServer procedures to an Fx application.
// It expects a MetaAPIYARPCServer to be present in the container.
//
//  fx.Provide(
//    apiv1.NewFxMetaAPIYARPCProcedures(),
//    ...
//  )
func NewFxMetaAPIYARPCProcedures() interface{} {
	return func(params FxMetaAPIYARPCProceduresParams) FxMetaAPIYARPCProceduresResult {
		return FxMetaAPIYARPCProceduresResult{
			Procedures: buildMetaAPIYARPCProcedures(buildMetaAPIYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "uber.cadence.api.v1.MetaAPI",
				FileDescriptors: yarpcFileDescriptorClosuref45a67f89288e7f3,
			},
		}
	}
}

type _MetaAPIYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_MetaAPIYARPCCaller) Health(ctx context.Context, request *HealthRequest, options ...yarpc.CallOption) (*HealthResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Health", request, newMetaAPIServiceHealthYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HealthResponse)
	if !ok {
		return nil, protobuf.CastError(emptyMetaAPIServiceHealthYARPCResponse, responseMessage)
	}
	return response, err
}

type _MetaAPIYARPCHandler struct {
	server MetaAPIYARPCServer
}

func (h *_MetaAPIYARPCHandler) Health(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *HealthRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*HealthRequest)
		if !ok {
			return nil, protobuf.CastError(emptyMetaAPIServiceHealthYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Health(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newMetaAPIServiceHealthYARPCRequest() proto.Message {
	return &HealthRequest{}
}

func newMetaAPIServiceHealthYARPCResponse() proto.Message {
	return &HealthResponse{}
}

var (
	emptyMetaAPIServiceHealthYARPCRequest  = &HealthRequest{}
	emptyMetaAPIServiceHealthYARPCResponse = &HealthResponse{}
)

var yarpcFileDescriptorClosuref45a67f89288e7f3 = [][]byte{
	// uber/cadence/api/v1/service_meta.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4b, 0x03, 0x41,
		0x10, 0x85, 0xd9, 0x2b, 0x92, 0x38, 0x60, 0x84, 0xb5, 0xf0, 0xb0, 0x0a, 0x27, 0x48, 0xaa, 0x59,
		0x4e, 0x3b, 0xad, 0x92, 0x4a, 0x0b, 0xe1, 0xbc, 0xc2, 0x42, 0x04, 0x99, 0x5b, 0x87, 0xcb, 0x12,
		0xf7, 0x76, 0xbd, 0xdd, 0xdb, 0xbf, 0x5f, 0xf2, 0xab, 0x08, 0x84, 0x94, 0xf3, 0xf8, 0x98, 0x37,
		0xf3, 0xc1, 0xfd, 0xd0, 0x70, 0xaf, 0x34, 0xfd, 0x70, 0xa7, 0x59, 0x91, 0x37, 0x2a, 0x95, 0x2a,
		0x70, 0x9f, 0x8c, 0xe6, 0x6f, 0xcb, 0x91, 0xd0, 0xf7, 0x2e, 0x3a, 0x79, 0xbd, 0xe1, 0x70, 0xcf,
		0x21, 0x79, 0x83, 0xa9, 0x2c, 0xae, 0xe0, 0xf2, 0x85, 0xe9, 0x37, 0xae, 0x6a, 0xfe, 0x1b, 0x38,
		0xc4, 0xe2, 0x09, 0xa6, 0x87, 0x20, 0x78, 0xd7, 0x05, 0x96, 0x53, 0xc8, 0xdc, 0x3a, 0x17, 0x33,
		0x31, 0x9f, 0xd4, 0x99, 0x5b, 0xcb, 0x1c, 0xc6, 0x96, 0x43, 0xa0, 0x96, 0xf3, 0x6c, 0x26, 0xe6,
		0x17, 0xf5, 0x61, 0x7c, 0xf8, 0x82, 0xf1, 0x1b, 0x47, 0x5a, 0x54, 0xaf, 0xf2, 0x1d, 0x46, 0xbb,
		0x35, 0xb2, 0xc0, 0x13, 0xbd, 0x78, 0x54, 0x7a, 0x7b, 0x77, 0x96, 0xd9, 0xdd, 0xb1, 0xfc, 0x80,
		0x1b, 0xed, 0xec, 0x29, 0x72, 0x39, 0x59, 0x78, 0x53, 0x6d, 0x9e, 0xac, 0xc4, 0xa7, 0x6a, 0x4d,
		0x5c, 0x0d, 0x0d, 0x6a, 0x67, 0xd5, 0x91, 0x19, 0x6c, 0xb9, 0x53, 0x5b, 0x15, 0x7b, 0x49, 0xcf,
		0xe4, 0x4d, 0x2a, 0x9b, 0xd1, 0x36, 0x7b, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x9d, 0x1b,
		0x11, 0x48, 0x01, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) MetaAPIYARPCClient {
			return NewMetaAPIYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
