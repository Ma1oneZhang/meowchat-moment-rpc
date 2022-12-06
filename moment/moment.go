// Code generated by goctl. DO NOT EDIT.
// Source: moment.proto

package moment

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddMomentReq      = pb.AddMomentReq
	AddMomentResp     = pb.AddMomentResp
	DeleteMomentReq   = pb.DeleteMomentReq
	DeleteMomentResp  = pb.DeleteMomentResp
	GetManyMomentReq  = pb.GetManyMomentReq
	GetManyMomentResp = pb.GetManyMomentResp
	GetMomentReq      = pb.GetMomentReq
	GetMomentResp     = pb.GetMomentResp
	Moment            = pb.Moment
	UpdateMomentReq   = pb.UpdateMomentReq
	UpdateMomentResp  = pb.UpdateMomentResp
	User              = pb.User

	Moment interface {
		GetManyMoment(ctx context.Context, in *GetManyMomentReq, opts ...grpc.CallOption) (*GetManyMomentResp, error)
		GetMoment(ctx context.Context, in *GetMomentReq, opts ...grpc.CallOption) (*GetMomentResp, error)
		AddMoment(ctx context.Context, in *AddMomentReq, opts ...grpc.CallOption) (*AddMomentResp, error)
		UpdateMoment(ctx context.Context, in *UpdateMomentReq, opts ...grpc.CallOption) (*UpdateMomentResp, error)
		DeleteMoment(ctx context.Context, in *DeleteMomentReq, opts ...grpc.CallOption) (*DeleteMomentResp, error)
	}

	defaultMoment struct {
		cli zrpc.Client
	}
)

func NewMoment(cli zrpc.Client) Moment {
	return &defaultMoment{
		cli: cli,
	}
}

func (m *defaultMoment) GetManyMoment(ctx context.Context, in *GetManyMomentReq, opts ...grpc.CallOption) (*GetManyMomentResp, error) {
	client := pb.NewMomentClient(m.cli.Conn())
	return client.GetManyMoment(ctx, in, opts...)
}

func (m *defaultMoment) GetMoment(ctx context.Context, in *GetMomentReq, opts ...grpc.CallOption) (*GetMomentResp, error) {
	client := pb.NewMomentClient(m.cli.Conn())
	return client.GetMoment(ctx, in, opts...)
}

func (m *defaultMoment) AddMoment(ctx context.Context, in *AddMomentReq, opts ...grpc.CallOption) (*AddMomentResp, error) {
	client := pb.NewMomentClient(m.cli.Conn())
	return client.AddMoment(ctx, in, opts...)
}

func (m *defaultMoment) UpdateMoment(ctx context.Context, in *UpdateMomentReq, opts ...grpc.CallOption) (*UpdateMomentResp, error) {
	client := pb.NewMomentClient(m.cli.Conn())
	return client.UpdateMoment(ctx, in, opts...)
}

func (m *defaultMoment) DeleteMoment(ctx context.Context, in *DeleteMomentReq, opts ...grpc.CallOption) (*DeleteMomentResp, error) {
	client := pb.NewMomentClient(m.cli.Conn())
	return client.DeleteMoment(ctx, in, opts...)
}
