package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-moment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMomentByMultiCommunityIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMomentByMultiCommunityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMomentByMultiCommunityIdLogic {
	return &SearchMomentByMultiCommunityIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMomentByMultiCommunityIdLogic) SearchMomentByMultiCommunityId(in *pb.SearchMomentByMultiCommunityIdReq) (*pb.ListMomentResp, error) {
	data, total, err := l.svcCtx.MomentModel.SearchByMultiCommunityId(l.ctx, in.CommunityIds, in.Keyword, in.Count, in.Skip)
	if err != nil {
		return nil, err
	}
	res := transformMomentSlice(data)
	return &pb.ListMomentResp{Moments: res, Total: total}, nil
}
