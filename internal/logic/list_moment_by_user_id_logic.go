package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMomentByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMomentByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMomentByUserIdLogic {
	return &ListMomentByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMomentByUserIdLogic) ListMomentByUserId(in *pb.ListMomentByUserIdReq) (*pb.ListMomentResp, error) {
	data, total, err := l.svcCtx.MomentModel.FindManyByUserId(l.ctx, in.UserId, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := transformMomentSlice(data)
	return &pb.ListMomentResp{Moments: res, Total: total}, nil
}
