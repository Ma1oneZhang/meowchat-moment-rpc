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
	res := make([]*pb.Moment, in.Count)
	for i, d := range data {
		m := &pb.Moment{
			Id:          d.ID.Hex(),
			CreateAt:    d.CreateAt.Unix(),
			Photos:      d.Photos,
			Title:       d.Title,
			Text:        d.Text,
			UserId:      d.UserId,
			CommunityId: d.CommunityId,
			CatId:       d.CatId,
		}
		res[i] = m
	}
	return &pb.ListMomentResp{Moments: res, Total: total}, nil
}
