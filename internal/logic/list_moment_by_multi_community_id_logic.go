package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMomentByMultiCommunityIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMomentByMultiCommunityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMomentByMultiCommunityIdLogic {
	return &ListMomentByMultiCommunityIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMomentByMultiCommunityIdLogic) ListMomentByMultiCommunityId(in *pb.ListMomentByMultiCommunityIdReq) (*pb.ListMomentResp, error) {
	data, total, err := l.svcCtx.MomentModel.FindManyByMultiCommunityId(l.ctx, in.CommunityIds, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := transformMomentSlice(data)
	return &pb.ListMomentResp{Moments: res, Total: total}, nil
}
